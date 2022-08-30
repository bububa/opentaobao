package tmallservice

// PagedResult 结构体
type PagedResult struct {
	// 核销单
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
	// 工单列表
	DataList []Workcard `json:"data_list,omitempty" xml:"data_list>workcard,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
