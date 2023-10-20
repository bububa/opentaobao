package alihouse

// AlibabaalihousenewhomereviewchangestatusResult 结构体
type AlibabaalihousenewhomereviewchangestatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 更新结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
