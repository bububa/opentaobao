package lstlogistics

// BaseResult 结构体
type BaseResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 发货单ID
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
}
