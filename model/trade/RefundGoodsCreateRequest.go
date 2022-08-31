package trade

// RefundGoodsCreateRequest 结构体
type RefundGoodsCreateRequest struct {
	// 退货商品列表
	RefundGoodsSubOrders []RefundGoodsSubOrder `json:"refund_goods_sub_orders,omitempty" xml:"refund_goods_sub_orders>refund_goods_sub_order,omitempty"`
	// 子订单号，默认传商品列表的第一个子订单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 门店标识
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 取货类型（&#34;FETCH_HOME&#34;：上门；&#34;ON_SHOP&#34;：到店；&#34;NONE&#34;：无需取）
	RefundFetchType string `json:"refund_fetch_type,omitempty" xml:"refund_fetch_type,omitempty"`
	// 主订单号
	MainBizOrderId string `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 买家标识
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 买家姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家联系方式
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 买家地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 小二备注
	OperatorMemo string `json:"operator_memo,omitempty" xml:"operator_memo,omitempty"`
	// 渠道来源，欧尚外卖默认是19
	InitFrom int64 `json:"init_from,omitempty" xml:"init_from,omitempty"`
}
