package jipiao

// ReturnApplyPassenge 结构体
type ReturnApplyPassenge struct {
	// 退款航段信息
	ReturnTicketSegment []ReturnTicketSegment `json:"return_ticket_segment,omitempty" xml:"return_ticket_segment>return_ticket_segment,omitempty"`
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 人ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 乘机人类型
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 退款手续费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款金额
	RefundMoney int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
	// 票价信息(分)
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 优惠后票面价
	DiscountTicketPrice int64 `json:"discount_ticket_price,omitempty" xml:"discount_ticket_price,omitempty"`
	// 优惠券金额
	VoucherPrice int64 `json:"voucher_price,omitempty" xml:"voucher_price,omitempty"`
}
