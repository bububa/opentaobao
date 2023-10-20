package alihouse

// AlibabaalihousenewhomeadvisermessagenoticeResult 结构体
type AlibabaalihousenewhomeadvisermessagenoticeResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或失败
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
