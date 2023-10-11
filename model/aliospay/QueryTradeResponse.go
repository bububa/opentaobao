package aliospay

// QueryTradeResponse 结构体
type QueryTradeResponse struct {
	// 支付结果状态,取值为:INIT初始，WAIT_BUYER_PAY : 等待用户付款。TRADE_SUCCESS:支付已经成功。 TRADE_CLOSED:未付款交易超时关闭，或支付完成后全额退款。TRADE_FINISHED交易结束，不可退款
	PayResult string `json:"pay_result,omitempty" xml:"pay_result,omitempty"`
	// 支付方式，SCAN-扫码，AUTH-免密支付，APP-手机支付，ALL_DEDUCTION-全额抵扣，SIGN_AND_PAY-签约并支付，WAP-手机网页支付，PERIOD_PAY-周期扣款
	PayWay string `json:"pay_way,omitempty" xml:"pay_way,omitempty"`
	// 斑马支付单Id
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 支付渠道，ALIPAY支付宝，WECHAT微信
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 支付渠道交易流水号
	ExtraTradeNo string `json:"extra_trade_no,omitempty" xml:"extra_trade_no,omitempty"`
	// 交易支付时间，未进行支付无值，时间戳
	PaymentTime int64 `json:"payment_time,omitempty" xml:"payment_time,omitempty"`
	// 实收金额，单位分
	ReceiptAmount int64 `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 订单总金额，单位分
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 订单创建时间，时间戳毫秒格式
	OrderTime int64 `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 退款时间，时间戳毫秒格式
	GmtRefund int64 `json:"gmt_refund,omitempty" xml:"gmt_refund,omitempty"`
	// 退款金额，单位为分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 补差金额，单位为分
	ReplenishAmount int64 `json:"replenish_amount,omitempty" xml:"replenish_amount,omitempty"`
}
