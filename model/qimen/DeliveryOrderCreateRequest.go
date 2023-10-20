package qimen

// DeliveryOrderCreateRequest 结构体
type DeliveryOrderCreateRequest struct {
	// 订单列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimendeliveryordercreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
