package perfect

// LoadReceiveResponse 结构体
type LoadReceiveResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 成功标记
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
