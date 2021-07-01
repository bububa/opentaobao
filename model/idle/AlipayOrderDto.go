package idle

// AlipayOrderDto 
type AlipayOrderDto struct {
    // 金额
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 支付状态，1未支付，6已支付
    PayStatus   int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    // 支付宝交易号
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
    // 支付订单
    PayOrderId   string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
}
