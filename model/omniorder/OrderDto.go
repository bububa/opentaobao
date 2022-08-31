package omniorder

// OrderDto 结构体
type OrderDto struct {
	// 子订单，属性与主订单相同
	DetailOrderList []string `json:"detail_order_list,omitempty" xml:"detail_order_list>string,omitempty"`
	// 消费者nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 旗舰店名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 1-父订单，0-子订单
	Main int64 `json:"main,omitempty" xml:"main,omitempty"`
	// 旗舰店id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 支付订单id
	PayOrderId int64 `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 支付状态：1 - 未冻结/未付款 -&gt;等待买家付款 2 - 已冻结/已付款 -&gt;等待卖家发货 4 - 已退款 -&gt;交易关闭 6 - 已转交易 -&gt; 交易成功 7 - 没有创建外部交易(支付宝交易) 8 - 交易被淘宝关闭 9 - 不可付款
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 物流状态：1 - 未发货 -&gt; 等待卖家发货, 2 - 已发货 -&gt; 等待买家确认收货, 3 - 已收货 -&gt; 交易成功, 4 - 已经退货 -&gt; 交易失败, 5 - 部分收货 -&gt; 交易成功, 6 - 部分发货中, 8 - 还未创建物流订单, 9 - 配货中,目前周期购已经使用到
	LogisticsStatus int64 `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 商品id
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 单件商品价格，单位分
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 商品的卖出金额调整，金额增加时为正数，金额减少时为负数，单位是分，不带小数。 如果是对整个订单的金额调整，会保存在主订单上，否则保存在子订单里
	AdjustFee int64 `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 系统对商品作的减价，始终为正数，单位是分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 退款状态：1 买家已经申请退款，等待卖家同意 2 卖家已经同意退款，等待买家退货 3 买家已经退货，等待卖家确认收货 4 退款关闭 5 退款成功 6 卖家拒绝退款 9 没有申请退款,对主定单和子订单都有效 10 有活动退款，申请退款后，只对主定单有效 只有一笔订单的情况，主定单的退款状态可能为1,2,3,6,10 11退款结束，只对父订单有效
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 到目前为止生效的退款，单位是分,父订单的退款等于各个子订单退款的和
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 实付金额，单位是分
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
}
