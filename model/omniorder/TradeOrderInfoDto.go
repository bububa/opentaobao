package omniorder

// TradeOrderInfoDto 结构体
type TradeOrderInfoDto struct {
	// 主订单ID
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 子订单信息
	SubOrders []TradeOrderDetailDto `json:"sub_orders,omitempty" xml:"sub_orders>trade_order_detail_dto,omitempty"`
}
