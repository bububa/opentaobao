package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"

	"github.com/bububa/opentaobao/metadata"
)

// Download 下载淘宝API文档metadata
// dir: 文件保存路径
// specificPkg: 指定API类目对应包名，参考package.json
func Download(dir string, specificPkg string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	clt := &http.Client{
		Jar: jar,
	}
	// 获取_tb_token_ cookie
	tbToken, err := getTbToken(clt)
	if err != nil {
		return err
	}

	// 获取文档目录树
	catelogTrees, err := getApiCatelogs(clt, tbToken)
	if err != nil {
		return err
	}

	pkgConfigMp := make(map[int64]metadata.PkgConfig, len(metadata.PkgsConfig))
	for _, cfg := range metadata.PkgsConfig {
		pkgConfigMp[cfg.Id] = cfg
	}

	// 遍历目录树
	for _, catelogTree := range catelogTrees {
		cfg, found := pkgConfigMp[catelogTree.Id]
		if !found { // package.json未定义分包
			log.Printf("[ERR] %d: %s, missing pkg config\n", catelogTree.Id, catelogTree.Name)
			for _, catelog := range catelogTree.CatelogList {
				log.Printf("[SKIP] %d: %s->%s\n", catelogTree.Id, catelogTree.Name, catelog.Name)
			}
			continue
		}

		// 如果设置了置顶API类目，跳过不相关类目
		if specificPkg != "" && cfg.Name != specificPkg {
			continue
		}

		// 保存单类目文档文件
		catelogPath, err := saveCatelog(cfg, dir)
		if err != nil {
			log.Printf("[ERR] %d: %s, %s\n", catelogTree.Id, catelogTree.Name, err.Error())
			continue
		}

		// 遍历API文档
		for _, catelog := range catelogTree.CatelogList {
			// 获取API文档
			doc, err := getDoc(clt, tbToken, catelog)
			if err != nil {
				log.Printf("[ERR] %d: %s->%s, %s\n", catelogTree.Id, catelogTree.Name, catelog.Name, err.Error())
				continue
			}
			doc.Id = catelog.Id // 下载API文档内缺少API ID
			// 保存API文档
			if err := saveDoc(doc, catelogPath); err != nil {
				log.Printf("[ERR] %d: %s->%s, %s\n", catelogTree.Id, catelogTree.Name, catelog.Name, err.Error())
				continue
			}
			log.Printf("[DONE] %d: %s->%s\n", catelogTree.Id, catelogTree.Name, catelog.Name)
		}
	}
	return nil
}

// getTbToken 获取_tb_token_ cookie
func getTbToken(clt *http.Client) (string, error) {
	resp, err := clt.Get(metadata.DOC_CENTER_URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "_tb_token_" {
			return cookie.Value, nil
		}
	}
	return "", errors.New("not found _tb_token_")
}

// getApiCatelogs 获取文档目录树
func getApiCatelogs(clt *http.Client, tbToken string) ([]metadata.ApiCatelogTree, error) {
	resp, err := clt.Get(fmt.Sprintf(metadata.API_CATELOG_URL, url.QueryEscape(tbToken)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ret metadata.ApiCatelogResponse
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, ret
	}
	if len(ret.Data.TreeCategories) == 0 {
		return nil, nil
	}
	return ret.Data.TreeCategories[0].CatelogTrees, nil
}

// getDoc 获取API文档
func getDoc(clt *http.Client, tbToken string, catelog metadata.ApiCatelog) (metadata.ApiDoc, error) {
	resp, err := clt.Get(fmt.Sprintf(metadata.DOC_URL, catelog.Id, url.QueryEscape(tbToken)))
	if err != nil {
		return metadata.ApiDoc{}, err
	}
	defer resp.Body.Close()
	var ret metadata.ApiDocResponse
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return metadata.ApiDoc{}, err
	}
	if ret.IsError() {
		return metadata.ApiDoc{}, ret
	}
	return ret.Data, nil
}

// saveCatelog 保存单类目文档文件
func saveCatelog(cfg metadata.PkgConfig, dir string) (string, error) {
	catelogPath := filepath.Join(dir, cfg.Pkg)
	log.Println(catelogPath)
	os.RemoveAll(catelogPath)
	if err := os.MkdirAll(catelogPath, os.ModePerm); err != nil {
		return catelogPath, err
	}
	buf, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return catelogPath, err
	}
	catelogTreeFilePath := filepath.Join(catelogPath, "catelog.json")
	if err := ioutil.WriteFile(catelogTreeFilePath, buf, os.ModePerm); err != nil {
		return catelogPath, err
	}
	return catelogPath, nil
}

// saveDoc 保存API文档文件
func saveDoc(doc metadata.ApiDoc, catelogPath string) error {
	buf, err := json.MarshalIndent(doc, "", "\t")
	if err != nil {
		return err
	}
	docPath := filepath.Join(catelogPath, fmt.Sprintf("%s.json", doc.Filename()))
	if err := ioutil.WriteFile(docPath, buf, os.ModePerm); err != nil {
		return err
	}
	return nil
}
