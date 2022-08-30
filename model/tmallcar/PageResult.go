package tmallcar

// PageResult 结构体
type PageResult struct {
	// 返回数据
	Models []TopOrderDto `json:"models,omitempty" xml:"models>top_order_dto,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 一共多少页 结果
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 分页大小
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页多少条记录
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 记录总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
