package alihouse

// AlibabaAlihouseNewhomeActivitySubscriptionBindResult 结构体
type AlibabaAlihouseNewhomeActivitySubscriptionBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否保存成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
