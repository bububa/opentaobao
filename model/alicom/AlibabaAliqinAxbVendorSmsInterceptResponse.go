package alicom

// AlibabaAliqinAxbVendorSmsInterceptResponse 结构体
type AlibabaAliqinAxbVendorSmsInterceptResponse struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 此字段忽略，只用判断code是否为OK
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// 响应的业务CODE：OK代表请求成功，非OK的错误码建议进行重试
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
