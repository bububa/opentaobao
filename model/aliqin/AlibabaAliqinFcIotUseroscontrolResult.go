package aliqin

// AlibabaaliqinfciotuseroscontrolResult 结构体
type AlibabaaliqinfciotuseroscontrolResult struct {
	// 停开结果描述
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 停开结果编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
