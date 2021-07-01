package aliospay

// PayRecordData 结构体
type PayRecordData struct {
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// alios支付订单id
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 支付宝订单id
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 支付结果状态,取值为:INIT初始，WAIT_BUYER_PAY : 等待用户付款。TRADE_SUCCESS:支付已经成功。 TRADE_CLOSED:未付款交易超时关闭，或支付完成后全额退款。TRADE_FINISHED交易结束，不可退款
	PayResult string `json:"pay_result,omitempty" xml:"pay_result,omitempty"`
	// 订单总金额，单位分
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 实收金额，单位分
	ReceiptAmount int64 `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 退款金额，单位分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 订单创建时间，时间戳
	OrderTime int64 `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 交易支付时间，未进行支付无值，时间戳
	PaymentTime int64 `json:"payment_time,omitempty" xml:"payment_time,omitempty"`
}
