package idle

// CommonResult 结构体
type CommonResult struct {
	// 异常码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 异常提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
