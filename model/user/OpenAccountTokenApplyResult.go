package user

// OpenAccountTokenApplyResult 结构体
type OpenAccountTokenApplyResult struct {
	// token
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
