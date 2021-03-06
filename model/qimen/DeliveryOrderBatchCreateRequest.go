package qimen

// DeliveryOrderBatchCreateRequest 结构体
type DeliveryOrderBatchCreateRequest struct {
	// 订单信息
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderBatchcreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
