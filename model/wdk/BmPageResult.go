package wdk

// BmPageResult 结构体
type BmPageResult struct {
	// 当前页码
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 对应data
	Data []PaiyangStatDataResult `json:"data,omitempty" xml:"data>paiyang_stat_data_result,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
