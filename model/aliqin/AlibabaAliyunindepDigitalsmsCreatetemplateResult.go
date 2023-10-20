package aliqin

// AlibabaaliyunindepdigitalsmscreatetemplateResult 结构体
type AlibabaaliyunindepdigitalsmscreatetemplateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 模板code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
