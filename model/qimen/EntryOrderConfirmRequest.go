package qimen

// EntryOrderConfirmRequest 结构体
type EntryOrderConfirmRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 分批次入库的最后一次回传
	TotalOrders []TotalOrder `json:"totalOrders,omitempty" xml:"totalOrders>total_order,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenentryorderconfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
