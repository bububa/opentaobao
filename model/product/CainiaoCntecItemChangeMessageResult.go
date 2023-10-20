package product

// CainiaocntecitemchangemessageResult 结构体
type CainiaocntecitemchangemessageResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功接受到请求
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
