package axintrade

// HotelOrderRefundResApiDto 结构体
type HotelOrderRefundResApiDto struct {
	// 退款单号
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}
