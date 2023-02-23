package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/bububa/opentaobao/metadata"
	"github.com/bububa/opentaobao/metadata/util"
)

// Gen 生成SDK入口
//
// metaPath: metadata 保存路径
// patchPath: patch 文件路径
// pkg: 指定API类目包名
func Gen(metaPath string, patchPath string, pkg string) error {
	rd, err := ioutil.ReadDir(metaPath)
	if err != nil {
		log.Fatalln(err)
	}

	pkgs := make([]ApiPkg, 0, len(rd)) // 生成的包，用于生成README文件

	// 遍历metadata目录
	for _, fi := range rd {
		if !fi.IsDir() {
			continue
		}

		// 如果指定API类目，跳过所有不匹配的类目
		if pkg != "" && fi.Name() != pkg {
			continue
		}

		catelogPath := filepath.Join(metaPath, fi.Name())       // 类目metadata路径
		catelogPatchPath := filepath.Join(patchPath, fi.Name()) // 类目patch文件路径

		// 生成API类目对应的api 和 model文件
		apiPkg := catelogHandler(catelogPath, catelogPatchPath)
		pkgs = append(pkgs, apiPkg)
	}

	// 如果未制定API类目则生成README文件
	if pkg == "" {
		err := genReadme(pkgs)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

// catelogHandler 生成API类目对应的api 和 model文件
func catelogHandler(catelogPath string, catelogPatchPath string) ApiPkg {
	// 获取类目包配置
	pkgCfg, err := getPkgConfig(catelogPath)

	apiPkg := ApiPkg{
		PkgConfig: pkgCfg,
	}

	if err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogPath, err.Error())
		return apiPkg
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogPath, err.Error())
		return apiPkg
	}

	// 生成类目包 model 路径
	catelogModelPath := filepath.Join(wd, util.StringsJoin("./model/", pkgCfg.Pkg))
	os.RemoveAll(catelogModelPath)
	if err := os.MkdirAll(catelogModelPath, os.ModePerm); err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogModelPath, err.Error())
		return apiPkg
	}

	// 生成类目包 api 路径
	catelogApiPath := filepath.Join(wd, util.StringsJoin("./api/", pkgCfg.Pkg))
	os.RemoveAll(catelogApiPath)
	if err := os.MkdirAll(catelogApiPath, os.ModePerm); err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogApiPath, err.Error())
		return apiPkg
	}
	log.Printf("[INFO] ID:%d, Name:%s, Pkg:%s\n", pkgCfg.Id, pkgCfg.Name, pkgCfg.Pkg)

	templatePath := filepath.Join(wd, "metadata/template") // SDK模版路径

	var models []metadata.TplModel // 包内结构体
	var files []string
	filesMp := make(map[string]struct{})
	catelogRd, err := ioutil.ReadDir(catelogPath)
	if err != nil {
		log.Fatalln(err)
	}
	// 遍历metadata文件
	for _, fi := range catelogRd {
		if fi.IsDir() {
			continue
		}
		if fi.Name() == "catelog.json" {
			continue
		}
		filesMp[fi.Name()] = struct{}{}
		files = append(files, fi.Name())
	}

	// 遍历patch目录文件, 用于弥补某些API metadata无法下载的情况
	if catelogPatchRd, err := ioutil.ReadDir(catelogPatchPath); err == nil {
		for _, fi := range catelogPatchRd {
			if fi.IsDir() {
				continue
			}
			if fi.Name() == "catelog.json" {
				continue
			}
			if _, found := filesMp[fi.Name()]; found {
				continue
			}
			files = append(files, fi.Name())
		}
	}

	for _, fi := range files {
		// 获取API文档
		doc, err := getDoc(catelogPath, catelogPatchPath, fi)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, File:%s, %s\n", pkgCfg.Name, fi, err.Error())
		}

		// 跳过API名为空的情况
		if doc.Name == "" {
			log.Printf("[ERR] Pkg:%s, File:%s, %+v\n", pkgCfg.Name, fi, doc)
			continue
		}

		// 设置返回结果的API包配置中的API文档链接
		if apiPkg.Link == "" {
			apiPkg.Link = metadata.DocLink(doc.Id)
		}

		// 转换API文档为SDK模版结构体
		tpl := doc.ApiTpl()
		tpl.Pkg = pkgCfg.Pkg // API模版结构体包名 = 包配置中的包名

		// 生成api包
		if err := genApi(templatePath, catelogApiPath, tpl); err != nil {
			log.Printf("[ERR] Pkg:%s, Api:%s, %s\n", pkgCfg.Name, doc.Name, err.Error())
			continue
		}

		// 生成model包
		apiModels, err := genApiModel(templatePath, catelogModelPath, tpl)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, Api:%s, %s\n", pkgCfg.Name, doc.Name, err.Error())
			continue
		}
		log.Printf("[INFO] Pkg:%s, Api:%s\n", pkgCfg.Name, doc.Name)
		models = append(models, apiModels...)
	}

	// 合并包内结构体
	models = metadata.MergeModels(models)

	for _, model := range models {
		model.Pkg = pkgCfg.Pkg
		// 生成结构体 go 文件
		if err := genModel(templatePath, catelogModelPath, model); err != nil {
			log.Printf("[ERR] Pkg:%s, Model:%s, %s\n", pkgCfg.Name, model.Name, err.Error())
			continue
		}
		log.Printf("[INFO] Pkg:%s, Model:%s\n", pkgCfg.Name, model.Name)
	}

	// 生成model doc.go
	if err := genPkgModelDoc(templatePath, catelogModelPath, apiPkg); err != nil {
		log.Printf("[ERR] Pkg:%s, Model doc.go, %s\n", pkgCfg.Name, err.Error())
	}

	// 生成api doc.go
	if err := genPkgApiDoc(templatePath, catelogApiPath, apiPkg); err != nil {
		log.Printf("[ERR] Pkg:%s, API doc.go, %s\n", pkgCfg.Name, err.Error())
	}
	return apiPkg
}

// getPkgConfig 获取API类目对应包设置
func getPkgConfig(catelogPath string) (metadata.PkgConfig, error) {
	var cfg metadata.PkgConfig
	fPath := filepath.Join(catelogPath, "catelog.json")
	pkgBytes, err := ioutil.ReadFile(fPath)
	if err != nil {
		return cfg, err
	}
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, pkgBytes); err != nil {
		return cfg, err
	}
	if err := json.NewDecoder(buf).Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

// getDoc 获取API文档
func getDoc(catelogPath string, catelogPatchPath string, filename string) (metadata.ApiDoc, error) {
	var doc metadata.ApiDoc

	// 先检查对应API文档是否有patch文件，如果有patch则使用patch代替
	fPath := filepath.Join(catelogPatchPath, filename)
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		fPath = filepath.Join(catelogPath, filename)
	} else {
		log.Printf("[WRN] doc:%s, use patch\n", fPath)
	}
	docBytes, err := ioutil.ReadFile(fPath)
	if err != nil {
		return doc, err
	}
	docBytes = util.PrintableBytes(docBytes)
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, docBytes); err != nil {
		return doc, err
	}
	if err := json.NewDecoder(buf).Decode(&doc); err != nil {
		return doc, err
	}
	return doc, nil
}

// genPkgApiDoc 生成api doc.go
func genPkgApiDoc(templatePath string, apiPath string, pkg ApiPkg) error {
	filePath := filepath.Join(templatePath, "pkg_api_doc.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(apiPath, "doc.go")
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, pkg)
	return gofmt(targetFile)
}

// genPkgModelDoc 生成model doc.go
func genPkgModelDoc(templatePath string, modelPath string, pkg ApiPkg) error {
	filePath := filepath.Join(templatePath, "pkg_model_doc.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, "doc.go")
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, pkg)
	return gofmt(targetFile)
}

// genApi 生成api文件
func genApi(templatePath string, apiPath string, tpl metadata.ApiTpl) error {
	filePath := filepath.Join(templatePath, "api.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(apiPath, fmt.Sprintf("%s.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, tpl)
	return gofmt(targetFile)
}

// genApiModel 生成model文件
// 返回 API 包含的对象结构体
func genApiModel(templatePath string, modelPath string, tpl metadata.ApiTpl) ([]metadata.TplModel, error) {
	// 生成APIRequest文件
	if err := genRequestModel(templatePath, modelPath, tpl); err != nil {
		return nil, err
	}
	// 生成APIResponse文件
	if err := genReponseModel(templatePath, modelPath, tpl); err != nil {
		return nil, err
	}
	reqModels := metadata.ExtractModels(tpl.RequestParams)   // APIRequest 中的结构体
	respModels := metadata.ExtractModels(tpl.ResponseParams) // APIResponse 中的结构体
	return metadata.MergeModels(append(reqModels, respModels...)), nil
}

// genRequestModel 生成APIRequest文件
func genRequestModel(templatePath string, modelPath string, tpl metadata.ApiTpl) error {
	filePath := filepath.Join(templatePath, "request.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%sAPIRequest.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	metadata.AlignTplParams(tpl.RequestParams)
	tmpl.Execute(fd, tpl)
	return gofmt(targetFile)
}

// genResponseModel 生成APIResponse文件
func genReponseModel(templatePath string, modelPath string, tpl metadata.ApiTpl) error {
	filePath := filepath.Join(templatePath, "response.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%sAPIResponse.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	metadata.AlignTplParams(tpl.ResponseParams)
	tmpl.Execute(fd, tpl)
	return gofmt(targetFile)
}

// genModel 生成对象结构体文件
func genModel(templatePath string, modelPath string, tpl metadata.TplModel) error {
	filePath := filepath.Join(templatePath, "model.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%s.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tpl.ImportModel = tpl.NeedImportModel()
	metadata.AlignTplParams(tpl.Params)
	tmpl.Execute(fd, tpl)
	return gofmt(targetFile)
}

// genReadme 生成README
func genReadme(pkgs []ApiPkg) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(wd, "metadata/template/readme.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(wd, "README.md")
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	slice := ApiPkgSlice(pkgs)
	sort.Sort(slice)
	tmpl.Execute(fd, map[string][]ApiPkg{
		"Pkgs": slice,
	})
	return nil
}

// gofmt 格式化go代码
func gofmt(target string) error {
	cmd := exec.Command("gofmt", "-s", "-w", target)
	return cmd.Run()
}
