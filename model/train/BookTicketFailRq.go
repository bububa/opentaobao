package train

// BookTicketFailRq 结构体
type BookTicketFailRq struct {
	// 出票失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
