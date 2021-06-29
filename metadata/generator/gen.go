package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/bububa/opentaobao/metadata"
)

func Gen(metaPath string, patchPath string, pkg string) error {
	rd, err := ioutil.ReadDir(metaPath)
	if err != nil {
		log.Fatalln(err)
	}
	var pkgs []ApiPkg
	for _, fi := range rd {
		if !fi.IsDir() {
			continue
		}
		if pkg != "" && fi.Name() != pkg {
			continue
		}
		catelogPath := filepath.Join(metaPath, fi.Name())
		catelogPatchPath := filepath.Join(patchPath, fi.Name())
		apiPkg := catelogHandler(catelogPath, catelogPatchPath)
		pkgs = append(pkgs, apiPkg)
	}
	if pkg == "" {
		err := genReadme(pkgs)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func catelogHandler(catelogPath string, catelogPatchPath string) ApiPkg {
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
	catelogModelPath := filepath.Join(wd, fmt.Sprintf("./model/%s", pkgCfg.Pkg))
	os.RemoveAll(catelogModelPath)
	if err := os.MkdirAll(catelogModelPath, os.ModePerm); err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogModelPath, err.Error())
		return apiPkg
	}
	catelogApiPath := filepath.Join(wd, fmt.Sprintf("./api/%s", pkgCfg.Pkg))
	os.RemoveAll(catelogApiPath)
	if err := os.MkdirAll(catelogApiPath, os.ModePerm); err != nil {
		log.Printf("[ERR] Path:%s, %s\n", catelogApiPath, err.Error())
		return apiPkg
	}
	log.Printf("[INFO] ID:%d, Name:%s, Pkg:%s\n", pkgCfg.Id, pkgCfg.Name, pkgCfg.Pkg)
	catelogRd, err := ioutil.ReadDir(catelogPath)
	if err != nil {
		log.Fatalln(err)
	}
	templatePath := filepath.Join(wd, "metadata/template")
	var models []metadata.TplModel
	var files []string
	filesMp := make(map[string]struct{})
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
		doc, err := getDoc(catelogPath, catelogPatchPath, fi)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, File:%s, %s\n", pkgCfg.Name, fi, err.Error())
		}

		if doc.Name == "" {
			log.Printf("[ERR] Pkg:%s, File:%s, %+v\n", pkgCfg.Name, fi, doc)
			continue
		}
		if apiPkg.Link == "" {
			apiPkg.Link = fmt.Sprintf(metadata.DOC_LINK, doc.Id)
		}
		// log.Printf("%+v\n", doc)
		// log.Printf("%+v\n", doc.ApiTpl())
		tpl := doc.ApiTpl()
		tpl.Pkg = pkgCfg.Pkg
		err = genApi(templatePath, catelogApiPath, tpl)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, Api:%s, %s\n", pkgCfg.Name, doc.Name, err.Error())
			continue
		}
		apiModels, err := genApiModel(templatePath, catelogModelPath, tpl)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, Api:%s, %s\n", pkgCfg.Name, doc.Name, err.Error())
			continue
		}
		log.Printf("[INFO] Pkg:%s, Api:%s\n", pkgCfg.Name, doc.Name)
		models = append(models, apiModels...)
	}
	models = metadata.MergeModels(models)
	for _, model := range models {
		model.Pkg = pkgCfg.Pkg
		err := genModel(templatePath, catelogModelPath, model)
		if err != nil {
			log.Printf("[ERR] Pkg:%s, Model:%s, %s\n", pkgCfg.Name, model.Name, err.Error())
			continue
		}
		log.Printf("[INFO] Pkg:%s, Model:%s\n", pkgCfg.Name, model.Name)
		if err := genPkgModelDoc(templatePath, catelogModelPath, apiPkg); err != nil {
			log.Printf("[ERR] Pkg:%s, Model doc.go, %s\n", pkgCfg.Name, err.Error())
		}
	}
	if err := genPkgApiDoc(templatePath, catelogApiPath, apiPkg); err != nil {
		log.Printf("[ERR] Pkg:%s, API doc.go, %s\n", pkgCfg.Name, err.Error())
	}
	return apiPkg
}

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

func getDoc(catelogPath string, catelogPatchPath string, filename string) (metadata.ApiDoc, error) {
	var doc metadata.ApiDoc
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
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, docBytes); err != nil {
		return doc, err
	}
	if err := json.NewDecoder(buf).Decode(&doc); err != nil {
		return doc, err
	}
	return doc, nil
}

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
	return nil
}

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
	return nil
}

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
	return nil
}

func genApiModel(templatePath string, modelPath string, tpl metadata.ApiTpl) ([]metadata.TplModel, error) {
	if err := genRequestModel(templatePath, modelPath, tpl); err != nil {
		return nil, err
	}
	if err := genReponseModel(templatePath, modelPath, tpl); err != nil {
		return nil, err
	}
	reqModels := metadata.ExtractModels(tpl.RequestParams)
	respModels := metadata.ExtractModels(tpl.ResponseParams)
	return metadata.MergeModels(append(reqModels, respModels...)), nil
}

func genRequestModel(templatePath string, modelPath string, tpl metadata.ApiTpl) error {
	filePath := filepath.Join(templatePath, "request.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%sRequest.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, tpl)
	return nil
}

func genReponseModel(templatePath string, modelPath string, tpl metadata.ApiTpl) error {
	filePath := filepath.Join(templatePath, "response.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%sResponse.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, tpl)
	return nil
}

func genModel(templatePath string, modelPath string, tpl metadata.TplModel) error {
	filePath := filepath.Join(templatePath, "model.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	tpl.ImportModel = tpl.NeedImportModel()
	targetFile := filepath.Join(modelPath, fmt.Sprintf("%s.go", tpl.Name))
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	tmpl.Execute(fd, tpl)
	return nil
}

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
