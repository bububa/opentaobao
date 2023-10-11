package ascp

// ResultWrapper 结构体
type ResultWrapper struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回信息
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
