package tmallhk

// ClearancePayOrderDO 
type ClearancePayOrderDO struct {
    // 支付宝买家ID
    AlipayBuyerId   string `json:"alipay_buyer_id,omitempty" xml:"alipay_buyer_id,omitempty"`
    // 支付单号
    PayOrderId   int64 `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
}
