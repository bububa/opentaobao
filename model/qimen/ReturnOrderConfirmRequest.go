package qimen

// ReturnOrderConfirmRequest 结构体
type ReturnOrderConfirmRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 退货单信息
	ReturnOrder *ReturnOrder `json:"returnOrder,omitempty" xml:"returnOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenReturnorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
