package product

// CainiaoCntecItemChangeMessageResult 结构体
type CainiaoCntecItemChangeMessageResult struct {
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功接受到请求
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
