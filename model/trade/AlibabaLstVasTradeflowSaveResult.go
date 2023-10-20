package trade

// AlibabaLstVasTradeflowSaveResult 结构体
type AlibabaLstVasTradeflowSaveResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
