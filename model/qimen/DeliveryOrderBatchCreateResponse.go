package qimen

// DeliveryOrderBatchCreateResponse 结构体
type DeliveryOrderBatchCreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单详情
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
}
