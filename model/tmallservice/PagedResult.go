package tmallservice

// PagedResult 结构体
type PagedResult struct {
	// 工单列表
	DataList []Workcard `json:"data_list,omitempty" xml:"data_list>workcard,omitempty"`
	// 总页数
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
