package alicom

// AlibabaAliqinAxbVendorPushCallEventResponse 结构体
type AlibabaAliqinAxbVendorPushCallEventResponse struct {
	// 接口调用成功 OK
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 事件接收成功 true
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
