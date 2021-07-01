package qimen

// DeliveryOrderBatchCreateAnswerRequest 结构体
type DeliveryOrderBatchCreateAnswerRequest struct {
	// 发货单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderBatchcreateAnswerMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
