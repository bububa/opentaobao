package alihouse

// AlibabaalihousenewhomeprojecttradeorderResult 结构体
type AlibabaalihousenewhomeprojecttradeorderResult struct {
	// 成功的楼盘Id
	Data []string `json:"data,omitempty" xml:"data>string,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
