package alicom

// AlibabaaliqinaxbvendorexceptionnosyncResponse 结构体
type AlibabaaliqinaxbvendorexceptionnosyncResponse struct {
	// 错误信息,OK代表受理成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码,OK代表受理成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// module
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
