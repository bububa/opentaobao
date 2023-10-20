package qimen

// ReturnOrderCreateRequest 结构体
type ReturnOrderCreateRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 退货单信息
	ReturnOrder *ReturnOrder `json:"returnOrder,omitempty" xml:"returnOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenreturnordercreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
