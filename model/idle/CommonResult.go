package idle

// CommonResult 结构体
type CommonResult struct {
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 异常描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 文件ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 请求结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
