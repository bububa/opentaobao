package btrip

// TrainRefundTicketInfo 结构体
type TrainRefundTicketInfo struct {
	// 退票票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 乘客ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 退票金额
	RefundFee float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 服务费退款
	RefundServiceFee float64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
}
