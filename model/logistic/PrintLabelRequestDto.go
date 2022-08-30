package logistic

// PrintLabelRequestDto 结构体
type PrintLabelRequestDto struct {
	// 物流订单号
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}
