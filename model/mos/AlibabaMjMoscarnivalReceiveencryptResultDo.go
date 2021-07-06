package mos

// AlibabaMjMoscarnivalReceiveencryptResultDo 结构体
type AlibabaMjMoscarnivalReceiveencryptResultDo struct {
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 总行数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 券结果
	Data *AlibabaMjMoscarnivalReceiveencryptData `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
