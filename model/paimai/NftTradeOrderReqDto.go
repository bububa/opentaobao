package paimai

// NftTradeOrderReqDto 结构体
type NftTradeOrderReqDto struct {
	// 订单ID多个用逗号分开
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
}
