package scs

// TaobaoOnebpDkxAccountAccountFreezeResultDto 结构体
type TaobaoOnebpDkxAccountAccountFreezeResultDto struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
