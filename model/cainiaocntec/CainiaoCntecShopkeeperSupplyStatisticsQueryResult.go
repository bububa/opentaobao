package cainiaocntec

// CainiaocntecshopkeepersupplystatisticsqueryResult 结构体
type CainiaocntecshopkeepersupplystatisticsqueryResult struct {
	// 门店统计结果
	ModelList []ActivityStatisticsDto `json:"model_list,omitempty" xml:"model_list>activity_statistics_dto,omitempty"`
	// trace_id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
