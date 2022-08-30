package ascp

// BindingConsignOrderGiftResult 结构体
type BindingConsignOrderGiftResult struct {
	// 绑赠列表
	GiftOrders []GiftOrder `json:"gift_orders,omitempty" xml:"gift_orders>gift_order,omitempty"`
	// 交易主单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 翱象发货单号
	ConsignOrderCode string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty"`
}
