package wdk

// PayChannel 结构体
type PayChannel struct {
	// 销售渠道（POS，APP);订单来源（淘宝、京东、三江购物网）。如：APP|淘宝
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 交易订单POS机器号
	PosNo string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// promotionCardInfo
	PromotionCardInfo string `json:"promotion_card_info,omitempty" xml:"promotion_card_info,omitempty"`
	// &#34;扩展属性，根据payType不同，可存放相关内容。MAP&lt;String,String&gt;的json化字符串子订单列表,key列表 &#34;btn&#34;  // 银行卡凭证号,bank_trade_no &#34;bmi&#34;  // 银行卡商户号,bank_merchant_id &#34;bti&#34;  // 银行卡终端号,bank_term_id &#34;brn&#34;  // 银行卡系统参考号,bank_ref_no &#34;damt&#34; // 优惠金额，discount_amount &#34;tno&#34;  // 支付工具号tool_no &#34;tsn&#34;  // 支付工具流水号，tool_serail_no&#34;
	PayAttributes string `json:"pay_attributes,omitempty" xml:"pay_attributes,omitempty"`
	// 口碑优惠信息，是一个json数组
	KoubeiCouponInfo string `json:"koubei_coupon_info,omitempty" xml:"koubei_coupon_info,omitempty"`
	// 支付渠道类型 1.默认 10-支付宝 20-微信支付 30-积分支付 40-储值卡支付 50-银行卡支付。有支付渠道的情况下，必填。 必须是翱象支持的支付渠道，否则报错。
	PayChannelType int64 `json:"pay_channel_type,omitempty" xml:"pay_channel_type,omitempty"`
	// 当前支付渠道的支付金额，单位分。有支付渠道的情况下，必填。
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 当前支付渠道的退款金额，单位分。有支付渠道的情况下，必填。
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 所属淘宝主订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 当前渠道付款金额(分)
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 当前渠道付款类型：0-付款码支付，1-盒马在线码－普通支付（pos），2-盒马在线码－免密协议支付（pos），3-银联刷卡（旧pos），4-现金（pos），5-支付宝离线码（pos），6-支付宝快捷支付（app），7-支付平台, 8-国际支付宝, 9-支付宝当面付, 10-网商银行信任付, 11-支付宝, 12-支付宝, 13-网商银行融易收, 14-现金, 15-银行卡, 16-支票, 17-三江购物券, 18-三江赊销
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 销售类型（正向销售1：逆向销售2。本接口应返回2）
	TrdType int64 `json:"trd_type,omitempty" xml:"trd_type,omitempty"`
	// 实付抹分金额(分)
	ActualPayAmount int64 `json:"actual_pay_amount,omitempty" xml:"actual_pay_amount,omitempty"`
	// adjustAmount
	AdjustAmount int64 `json:"adjust_amount,omitempty" xml:"adjust_amount,omitempty"`
	// 支付宝优惠金额，单位分
	PmtAlipayPromotionAmt int64 `json:"pmt_alipay_promotion_amt,omitempty" xml:"pmt_alipay_promotion_amt,omitempty"`
	// 支付宝随机立减金额，单位分
	PmtAlipayReduceAmt int64 `json:"pmt_alipay_reduce_amt,omitempty" xml:"pmt_alipay_reduce_amt,omitempty"`
	// 口碑券的优惠金额，单位分
	PmtKoubeiCouponAmt int64 `json:"pmt_koubei_coupon_amt,omitempty" xml:"pmt_koubei_coupon_amt,omitempty"`
}
