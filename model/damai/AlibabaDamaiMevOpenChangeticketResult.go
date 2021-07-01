package damai

// AlibabaDamaiMevOpenChangeticketResult 结构体
type AlibabaDamaiMevOpenChangeticketResult struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回内容
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
