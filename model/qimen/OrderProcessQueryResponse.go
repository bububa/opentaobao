package qimen

// OrderProcessQueryResponse 结构体
type OrderProcessQueryResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单处理流程
	OrderProcess *OrderProcess `json:"orderProcess,omitempty" xml:"orderProcess,omitempty"`
}
