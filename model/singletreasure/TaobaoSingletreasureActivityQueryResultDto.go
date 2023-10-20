package singletreasure

// TaobaoSingletreasureActivityQueryResultDto 结构体
type TaobaoSingletreasureActivityQueryResultDto struct {
	// 查询结果
	DataList []ActivityInfo `json:"data_list,omitempty" xml:"data_list>activity_info,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 页码
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回结果个数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 系统执行是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
