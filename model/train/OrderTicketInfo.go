package train

// OrderTicketInfo 结构体
type OrderTicketInfo struct {
	// 出票结果定制信息列表
	VipCustomResultList []VipCustomTicketConfirmItem `json:"vip_custom_result_list,omitempty" xml:"vip_custom_result_list>vip_custom_ticket_confirm_item,omitempty"`
	// 座位号
	SeatNum string `json:"seat_num,omitempty" xml:"seat_num,omitempty"`
	// 车次
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 证件类型
	CertType string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 证件号
	CertificateNum string `json:"certificate_num,omitempty" xml:"certificate_num,omitempty"`
	// 12306票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// ttp子单ID
	TtpSubOrderId int64 `json:"ttp_sub_order_id,omitempty" xml:"ttp_sub_order_id,omitempty"`
	// 真实票价
	RealTicketPrice int64 `json:"real_ticket_price,omitempty" xml:"real_ticket_price,omitempty"`
	// 真实坐席
	RealSeat int64 `json:"real_seat,omitempty" xml:"real_seat,omitempty"`
	// 保险支付金额
	InsurancePayPrice int64 `json:"insurance_pay_price,omitempty" xml:"insurance_pay_price,omitempty"`
	// 定制票出票结果 1:定制票出票 0:非定制票出票
	VipCustomResult int64 `json:"vip_custom_result,omitempty" xml:"vip_custom_result,omitempty"`
}
