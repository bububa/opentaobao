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

func Download(dir string, specificPkg string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	clt := &http.Client{
		Jar: jar,
	}
	tbToken, err := getTbToken(clt)
	if err != nil {
		return err
	}
	catelogTrees, err := getApiCatelogs(clt, tbToken)
	if err != nil {
		return err
	}
	pkgConfigMp := make(map[int64]metadata.PkgConfig, len(metadata.PkgsConfig))
	for _, cfg := range metadata.PkgsConfig {
		pkgConfigMp[cfg.Id] = cfg
	}
	for _, catelogTree := range catelogTrees {
		cfg, found := pkgConfigMp[catelogTree.Id]
		if !found {
			log.Printf("[ERR] %d: %s, missing pkg config\n", catelogTree.Id, catelogTree.Name)
			for _, catelog := range catelogTree.CatelogList {
				log.Printf("[SKIP] %d: %s->%s\n", catelogTree.Id, catelogTree.Name, catelog.Name)
			}
			continue
		}
		if specificPkg != "" && cfg.Name != specificPkg {
			continue
		}

		catelogPath, err := saveCatelog(cfg, dir)
		if err != nil {
			log.Printf("[ERR] %d: %s, %s\n", catelogTree.Id, catelogTree.Name, err.Error())
			continue
		}
		for _, catelog := range catelogTree.CatelogList {
			doc, err := getDoc(clt, tbToken, catelog)
			if err != nil {
				log.Printf("[ERR] %d: %s->%s, %s\n", catelogTree.Id, catelogTree.Name, catelog.Name, err.Error())
				continue
			}
			doc.Id = catelog.Id
			if err := saveDoc(doc, catelogPath); err != nil {
				log.Printf("[ERR] %d: %s->%s, %s\n", catelogTree.Id, catelogTree.Name, catelog.Name, err.Error())
				continue
			}
			log.Printf("[DONE] %d: %s->%s\n", catelogTree.Id, catelogTree.Name, catelog.Name)
		}
	}
	return nil
}

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
