package alsc

// BaseResponse 结构体
type BaseResponse struct {
	// 链路ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误原因
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 拓展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 返回数据
	ResultObj *EntityPrizeTokenResp `json:"result_obj,omitempty" xml:"result_obj,omitempty"`
	// 可重试
	CanRetry bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
