package omniorder

// TaobaoOmniDealerOdersGetData 结构体
type TaobaoOmniDealerOdersGetData struct {
	// 子订单
	SubOrders []SubOrders `json:"sub_orders,omitempty" xml:"sub_orders>sub_orders,omitempty"`
	// 订单创建时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 买家呢称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 省
	ReceiverProv string `json:"receiver_prov,omitempty" xml:"receiver_prov,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 退款状态；可选值:WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意), WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货),  WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货),  SELLER_REFUSE_BUYER(卖家拒绝退款),  CLOSED(退款关闭),  SUCCESS(退款成功)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 多级时效预约单结束预约时间
	ExpectEnd string `json:"expect_end,omitempty" xml:"expect_end,omitempty"`
	// 履约门店的商家门店编码
	FulfillmentStoreOutId string `json:"fulfillment_store_out_id,omitempty" xml:"fulfillment_store_out_id,omitempty"`
	// 门店履约类型，可选值：PICKED_UP_IN_STORES(门店自提), STORE_DELIVERY(门店发货)
	FulfillmentType string `json:"fulfillment_type,omitempty" xml:"fulfillment_type,omitempty"`
	// 市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 发货时间
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 收货人的电话号码
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 镇
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
	// 卖家呢称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 收件人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 手机
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 区/县
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
	// 时效类型，可选值：NONE(无时效), APPOINTMENT_AGING(预约配送), HOUR_AGING(小时达), DAY_0_AGING(当日达), DAY_1_AGING(次日达), DAY_2_AGING(隔日达)
	AgingType string `json:"aging_type,omitempty" xml:"aging_type,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 收货地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 多级时效预约单开始预约时间
	ExpectStart string `json:"expect_start,omitempty" xml:"expect_start,omitempty"`
	// 卖家备注
	SellerMemo string `json:"seller_memo,omitempty" xml:"seller_memo,omitempty"`
	// 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 买家备注
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 订单状态，可选值：WAIT_BUYER_PAY(等待买家付款),  WAIT_SELLER_SEND_GOODS(等待卖家发货),  SELLER_CONSIGNED_PART(卖家部分发货),  WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货]),  TRADE_BUYER_SIGNED(买家已签收（货到付款专用）),  TRADE_FINISHED(交易成功),  TRADE_CLOSED(交易关闭),  TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭),  TRADE_NO_CREATE_PAY(没有创建外部交易（支付宝交易）)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 优惠金额，单位：分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 履约门店ID
	FulfillmentStoreId int64 `json:"fulfillment_store_id,omitempty" xml:"fulfillment_store_id,omitempty"`
	// 主订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 买家实际支付给卖家的金额，单位：分
	ActualTotalFee int64 `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 手工调整金额，单位：分
	AdjustFee int64 `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 经销商卖家ID
	DealerSellerId int64 `json:"dealer_seller_id,omitempty" xml:"dealer_seller_id,omitempty"`
	// 邮费单位分
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 购物金信息输出
	ExpandCardInfo *ExpandCardInfo `json:"expand_card_info,omitempty" xml:"expand_card_info,omitempty"`
	// 总金额，单价×数量，不包含运费和折扣，单位：分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
}
