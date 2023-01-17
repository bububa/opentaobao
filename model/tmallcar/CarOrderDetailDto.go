package tmallcar

// CarOrderDetailDto 结构体
type CarOrderDetailDto struct {
	// 子订单列表
	SubOrders []CarSubOrderDetailDto `json:"sub_orders,omitempty" xml:"sub_orders>car_sub_order_detail_dto,omitempty"`
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 用户填写购车人信息
	BuyerNameCollect string `json:"buyer_name_collect,omitempty" xml:"buyer_name_collect,omitempty"`
	// 用户填写购车身份证信息
	BuyerIdentityCollect string `json:"buyer_identity_collect,omitempty" xml:"buyer_identity_collect,omitempty"`
	// 用户填写手机号
	BuyerMobile string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
	// 交付服务地址
	CarDelivery string `json:"car_delivery,omitempty" xml:"car_delivery,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 实付金额，单位分
	ActualPaidFee int64 `json:"actual_paid_fee,omitempty" xml:"actual_paid_fee,omitempty"`
	// 退款状态：1买家申请退款，等待卖家同意；2卖家已同意退款，等待买家退货；3买家已退货，等待卖家确认收货；4退款关闭；5退款成功；6卖家拒绝退款；9没有申请退款
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 支付状态：1未付款；2已付款；4已退款(交易关闭)；6交易成功；8未支付订单关闭
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付单
	PayOrder *CarPayOrderDto `json:"pay_order,omitempty" xml:"pay_order,omitempty"`
	// 订单sku信息
	Sku *TradeItemSkuDto `json:"sku,omitempty" xml:"sku,omitempty"`
	// 订单商品信息（子订单有效）
	Item *TradeItemDto `json:"item,omitempty" xml:"item,omitempty"`
}
