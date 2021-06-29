package mos

// RefundResponse 
type RefundResponse struct {
    // 喵街交易凭证号。必然返回
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    // 原支付请求的商户订单号。必然返回
    OutTradeNo   string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
    // 外部退款流水号
    OutRequestNo   string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
    // 退款状态。FAIL退款失败，REFUNDING表示退款请求中，SUCCESS退款成功
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 买家在支付宝的用户id
    BuyerAlipayId   string `json:"buyer_alipay_id,omitempty" xml:"buyer_alipay_id,omitempty"`
    // 总退款金额。单位为人民币（分）
    RefundAmount   int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
    // 退款资金渠道列表
    FundBillList   []FundBillDO `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill_do,omitempty"`
    // 授权码来源。MJ：喵街，M_TAO：手淘，ALIPAY：支付宝
    AuthCodeSource   string `json:"auth_code_source,omitempty" xml:"auth_code_source,omitempty"`
    // 消费者喵街昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
}
