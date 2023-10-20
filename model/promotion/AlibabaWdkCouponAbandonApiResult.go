package promotion

// AlibabawdkcouponabandonApiResult 结构体
type AlibabawdkcouponabandonApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 操作结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
