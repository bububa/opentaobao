package openmall

// TopOrderVo 结构体
type TopOrderVo struct {
	// 子订单发货时间，当卖家对订单进行了多次发货，子订单的发货时间和主订单的发货时间可能不一样了，那么就需要以子订单的时间为准。（没有进行多次发货的订单，主订单的发货时间和子订单的发货时间都一样）
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 子订单的交易结束时间说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 子订单预计发货时间
	EstimateConTime string `json:"estimate_con_time,omitempty" xml:"estimate_con_time,omitempty"`
	// 商品备注
	ItemMemo string `json:"item_memo,omitempty" xml:"item_memo,omitempty"`
	// 购买数量。取值范围:大于零的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 商品数字ID
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 子订单编号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 外部网店自己定义的Sku编号
	OuterSkuId string `json:"outer_sku_id,omitempty" xml:"outer_sku_id,omitempty"`
	// 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。对于多子订单的交易，计算公式如下：payment = price * num + adjust_fee - discount_fee ；单子订单交易，payment与主订单的payment一致，对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 最近退款ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款状态。退款状态。可选值 WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// SKU的值。如：机身颜色:黑色;手机套餐:官方标配
	SkuPropertiesName string `json:"sku_properties_name,omitempty" xml:"sku_properties_name,omitempty"`
	// 订单状态。可选值: TRADE_NO_CREATE_PAY(没有创建支付宝交易，暂无) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用，暂无) TRADE_FINISHED(交易成功) TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易) PAY_PENDING(国际信用卡支付付款确认中，暂无)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单级订单优惠金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 是否发货
	IsShShip bool `json:"is_sh_ship,omitempty" xml:"is_sh_ship,omitempty"`
	// 商品SKUID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 物流公司名称
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 子订单包裹运输号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
}
