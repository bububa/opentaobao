package tblogistics

// WriteOffOrderDto 结构体
type WriteOffOrderDto struct {
	// 交易单所包含的商品列表
	GoodsList []WriteOffGoodsDto `json:"goods_list,omitempty" xml:"goods_list>write_off_goods_dto,omitempty"`
	// 核销订单Id
	LpOrderId string `json:"lp_order_id,omitempty" xml:"lp_order_id,omitempty"`
	// 淘宝交易Id
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
}
