package btrip

// FlightRefundTicketInfo 结构体
type FlightRefundTicketInfo struct {
	// 退票票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 退票原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 退票单id
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 退票类型：0自愿/1非自愿
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// 退票金额
	RefundTicketFee float64 `json:"refund_ticket_fee,omitempty" xml:"refund_ticket_fee,omitempty"`
}
