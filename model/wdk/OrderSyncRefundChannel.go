package wdk

// OrderSyncRefundChannel 结构体
type OrderSyncRefundChannel struct {
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款渠道
	RefundChannel int64 `json:"refund_channel,omitempty" xml:"refund_channel,omitempty"`
	// 退款单id
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}
