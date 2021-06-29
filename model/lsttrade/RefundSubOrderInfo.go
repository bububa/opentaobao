package lsttrade

// RefundSubOrderInfo 
type RefundSubOrderInfo struct {
    // 退款数量
    RefundCount   int64 `json:"refund_count,omitempty" xml:"refund_count,omitempty"`
    // 退款金额，单位分
    RefundPayment   int64 `json:"refund_payment,omitempty" xml:"refund_payment,omitempty"`
    // 子订单id
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}
