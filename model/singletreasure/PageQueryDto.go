package singletreasure

// PageQueryDto 结构体
type PageQueryDto struct {
	// 查询条件请求体
	Query *ActivityInfoListQueryDto `json:"query,omitempty" xml:"query,omitempty"`
	// 页码,最大为50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页数，从1开始
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
}
