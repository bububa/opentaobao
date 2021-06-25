package metadata

import "strings"

type ApiCatelogResponse struct {
	Success bool           `json:"success,omitempty"`
	Code    string         `json:"code,omitempty"`
	ErrMsg  string         `json:"errMsg,omitempty"`
	Data    ApiCatelogData `json:"data,omitempty"`
}

func (res ApiCatelogResponse) IsError() bool {
	return !res.Success
}

func (res ApiCatelogResponse) Error() string {
	return res.ErrMsg
}

type ApiCatelogData struct {
	Id             int64             `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	TreeCategories []ApiTreeCategory `json:"treeCategories,omitempty"`
}

type ApiTreeCategory struct {
	Id           int64            `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	CatelogTrees []ApiCatelogTree `json:"catelogTrees,omitempty"`
}

type ApiCatelogTree struct {
	Id          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	CatelogList []ApiCatelog `json:"catelogList,omitempty"`
}

type ApiCatelog struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	DocType int    `json:"docType,omitempty"`
	DocId   int64  `json:"docId,omitempty"`
	Tips    string `json:"tips,omitempty"`
	SubName string `json:"subName,omitempty"`
}

func (c ApiCatelog) Title() string {
	name := strings.Title(c.Name)
	return strings.ReplaceAll(name, ".", "")
}
