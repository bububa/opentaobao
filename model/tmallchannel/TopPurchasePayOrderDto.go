package tmallchannel

// TopPurchasePayOrderDto 
type TopPurchasePayOrderDto struct {
    // 支付流水编号
    PayOrderId   string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
    // 支付金额
    PayFee   int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}
