package aliqin

// AlibabaaliqinfciotdeviceisexistResult 结构体
type AlibabaaliqinfciotdeviceisexistResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否存在
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
