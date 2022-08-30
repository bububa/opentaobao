package idleisv

// CommonResult 结构体
type CommonResult struct {
	// 属性list
	Data []YoupinCpvResultDto `json:"data,omitempty" xml:"data>youpin_cpv_result_dto,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
