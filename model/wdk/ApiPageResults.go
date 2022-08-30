package wdk

// ApiPageResults 结构体
type ApiPageResults struct {
	// 业务结果集
	Model []BillDailyDto `json:"model,omitempty" xml:"model>bill_daily_dto,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 当前页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
