package qimen

// StockOutCreateRequest 结构体
type StockOutCreateRequest struct {
	// 单据信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 出库单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockoutCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
