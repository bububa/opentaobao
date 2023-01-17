package train

// ChangeApplySubOrderDto 结构体
type ChangeApplySubOrderDto struct {
	// 乘车人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 原票坐席号
	OriginalSeatNo string `json:"original_seat_no,omitempty" xml:"original_seat_no,omitempty"`
	// 原票坐席类型code
	OriginalSeatTypeCode string `json:"original_seat_type_code,omitempty" xml:"original_seat_type_code,omitempty"`
	// 原票坐席类型名称
	OriginalSeatTypeName string `json:"original_seat_type_name,omitempty" xml:"original_seat_type_name,omitempty"`
	// 证件号
	CertificateNo string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	// 证件类型名称
	CertificateTypeName string `json:"certificate_type_name,omitempty" xml:"certificate_type_name,omitempty"`
	// 证件类型code
	CertificateTypeCode string `json:"certificate_type_code,omitempty" xml:"certificate_type_code,omitempty"`
	// 原票票价(分)
	OriginalTicketPrice string `json:"original_ticket_price,omitempty" xml:"original_ticket_price,omitempty"`
	// 原票车厢名称
	OriginalCoachName string `json:"original_coach_name,omitempty" xml:"original_coach_name,omitempty"`
	// 票类型名称
	TicketTypeName string `json:"ticket_type_name,omitempty" xml:"ticket_type_name,omitempty"`
	// 原票车厢号
	OriginalCoachNo string `json:"original_coach_no,omitempty" xml:"original_coach_no,omitempty"`
	// 改签票坐席类型code
	ChangeSeatTypeCode string `json:"change_seat_type_code,omitempty" xml:"change_seat_type_code,omitempty"`
	// 改签票坐席类型名称
	ChangeSeatTypeName string `json:"change_seat_type_name,omitempty" xml:"change_seat_type_name,omitempty"`
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 改签申请单id
	ChangeApplyId int64 `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// 子单单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 票类型code
	TicketTypeCode int64 `json:"ticket_type_code,omitempty" xml:"ticket_type_code,omitempty"`
	// 用户提交改签票价(卧铺、儿童票都是按最高票价提交)
	PreTicketPrice int64 `json:"pre_ticket_price,omitempty" xml:"pre_ticket_price,omitempty"`
}
