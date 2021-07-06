package qimen

// EntryOrderConfirmRequest 结构体
type EntryOrderConfirmRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenEntryorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
