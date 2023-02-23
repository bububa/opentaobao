package paimai

// NftTradeOrderDto 结构体
type NftTradeOrderDto struct {
	// 订单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
