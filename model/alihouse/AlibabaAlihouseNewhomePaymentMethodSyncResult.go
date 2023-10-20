package alihouse

// AlibabaAlihouseNewhomePaymentMethodSyncResult 结构体
type AlibabaAlihouseNewhomePaymentMethodSyncResult struct {
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
