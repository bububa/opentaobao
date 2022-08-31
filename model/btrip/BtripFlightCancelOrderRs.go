package btrip

// BtripFlightCancelOrderRs 结构体
type BtripFlightCancelOrderRs struct {
	// 取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 失败类型
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
}
