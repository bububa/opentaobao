package mos

// OnsiteRefundResponse 结构体
type OnsiteRefundResponse struct {
	// fundBillList
	FundBillList []FundBillDo `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill_do,omitempty"`
	// 订单号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 外部订单号
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 外部退款流水号
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 退款状态。FAIL退款失败，REFUNDING表示退款请求中，SUCCESS退款成功
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// buyerAlipayId
	BuyerAlipayId string `json:"buyer_alipay_id,omitempty" xml:"buyer_alipay_id,omitempty"`
	// 授权码来源。MJ：喵街，M_TAO：手淘，ALIPAY：支付宝
	AuthCodeSource string `json:"auth_code_source,omitempty" xml:"auth_code_source,omitempty"`
	// 消费者喵街昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 本次退款请求，对应的退款金额（分）
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}
