package car

// OrderDetailInfo 结构体
type OrderDetailInfo struct {
	// 关单原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// createdTime
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 交易结束时间,确认时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// modifiedTime
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// orderSource
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// payTime
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// payTimeOutTime
	PayTimeOutTime string `json:"pay_time_out_time,omitempty" xml:"pay_time_out_time,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 折扣费用(优惠金额)
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// orderId
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// payStatus
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// realPay
	RealPay int64 `json:"real_pay,omitempty" xml:"real_pay,omitempty"`
	// refundFee
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// refundStatus
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// version
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
