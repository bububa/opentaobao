package metadata

// 下载淘宝API文档类目返回结果结构体
type ApiCatelogResponse struct {
	Success bool           `json:"success,omitempty"`
	Code    string         `json:"code,omitempty"`
	ErrMsg  string         `json:"errMsg,omitempty"`
	Data    ApiCatelogData `json:"data,omitempty"`
}

func (res ApiCatelogResponse) IsError() bool {
	return !res.Success
}

// implement Error interface
func (res ApiCatelogResponse) Error() string {
	return res.ErrMsg
}

// 淘宝API文档类目结果结构体
type ApiCatelogData struct {
	Id             int64             `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	TreeCategories []ApiTreeCategory `json:"treeCategories,omitempty"`
}

// 淘宝API文档类目树结构体
type ApiTreeCategory struct {
	Id           int64            `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	CatelogTrees []ApiCatelogTree `json:"catelogTrees,omitempty"`
}

// 淘宝API文档单类目结构体
type ApiCatelogTree struct {
	Id          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	CatelogList []ApiCatelog `json:"catelogList,omitempty"`
}

// 淘宝API文档单API结构体
type ApiCatelog struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	DocType int    `json:"docType,omitempty"`
	DocId   int64  `json:"docId,omitempty"`
	Tips    string `json:"tips,omitempty"`
	SubName string `json:"subName,omitempty"`
}
