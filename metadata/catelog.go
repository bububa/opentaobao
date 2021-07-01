package metadata

// ApiCatelogResponse 载淘宝API文档类目返回结果结构体
type ApiCatelogResponse struct {
	Success bool           `json:"success,omitempty"` // 是否成功
	Code    string         `json:"code,omitempty"`    // 错误代码
	ErrMsg  string         `json:"errMsg,omitempty"`  // 错误信息
	Data    ApiCatelogData `json:"data,omitempty"`    // 返回成功数据
}

// IsError 判断是否为错误
func (res ApiCatelogResponse) IsError() bool {
	return !res.Success
}

// Error implement Error interface
func (res ApiCatelogResponse) Error() string {
	return res.ErrMsg
}

// ApiCatelogData 淘宝API文档类目结果结构体
type ApiCatelogData struct {
	Id             int64             `json:"id,omitempty"`             // ID
	Name           string            `json:"name,omitempty"`           // 名称
	TreeCategories []ApiTreeCategory `json:"treeCategories,omitempty"` // 类目树列表
}

// ApiTreeCatelog 淘宝API文档类目树结构体
type ApiTreeCategory struct {
	Id           int64            `json:"id,omitempty"`           // api catelog tree id
	Name         string           `json:"name,omitempty"`         // api catelog tree name
	CatelogTrees []ApiCatelogTree `json:"catelogTrees,omitempty"` // 类目列表
}

// ApiCatelogTree 淘宝API文档单类目结构体
type ApiCatelogTree struct {
	Id          int64        `json:"id,omitempty"`          // 类目ID
	Name        string       `json:"name,omitempty"`        // 类目名
	CatelogList []ApiCatelog `json:"catelogList,omitempty"` // API列表
}

// ApiCatelog 淘宝API文档单API结构体
type ApiCatelog struct {
	Id      int64  `json:"id,omitempty"`      // API ID
	Name    string `json:"name,omitempty"`    // API 方法名
	DocType int    `json:"docType,omitempty"` // API 文档类型
	DocId   int64  `json:"docId,omitempty"`   // API 文档ID
	Tips    string `json:"tips,omitempty"`    // tips
	SubName string `json:"subName,omitempty"` // subname
}
