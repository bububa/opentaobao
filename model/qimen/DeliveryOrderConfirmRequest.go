package qimen

// DeliveryOrderConfirmRequest 结构体
type DeliveryOrderConfirmRequest struct {
	// 包裹信息
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// 单据列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
