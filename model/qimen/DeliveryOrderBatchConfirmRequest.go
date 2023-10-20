package qimen

// DeliveryOrderBatchConfirmRequest 结构体
type DeliveryOrderBatchConfirmRequest struct {
	// 发货单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimendeliveryorderbatchconfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
