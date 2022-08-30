package wdk

// PageInfoDto 结构体
type PageInfoDto struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
