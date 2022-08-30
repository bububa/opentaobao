package omniorder

// SubOrders 结构体
type SubOrders struct {
	// 子单修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 退款状态；可选值:WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意), WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货),  WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货),  SELLER_REFUSE_BUYER(卖家拒绝退款),  CLOSED(退款关闭),  SUCCESS(退款成功)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 子单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 经销商的货品ID
	DealerScItemId string `json:"dealer_sc_item_id,omitempty" xml:"dealer_sc_item_id,omitempty"`
	// 商品编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 订单状态，可选值：WAIT_BUYER_PAY(等待买家付款),  WAIT_SELLER_SEND_GOODS(等待卖家发货),  SELLER_CONSIGNED_PART(卖家部分发货),  WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货]),  TRADE_BUYER_SIGNED(买家已签收（货到付款专用）),  TRADE_FINISHED(交易成功),  TRADE_CLOSED(交易关闭),  TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭),  TRADE_NO_CREATE_PAY(没有创建外部交易（支付宝交易）)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品购买数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 子订单ID
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 商品价格，单位：分
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 实收金额，单位：分
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
