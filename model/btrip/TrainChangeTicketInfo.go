package btrip

// TrainChangeTicketInfo 结构体
type TrainChangeTicketInfo struct {
	// 改签票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 改签原票
	OriginTicketNo string `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	// 改签车厢号
	ChangeCoachNo string `json:"change_coach_no,omitempty" xml:"change_coach_no,omitempty"`
	// 改签座位号
	ChangeSeatNo string `json:"change_seat_no,omitempty" xml:"change_seat_no,omitempty"`
	// 改签车次类型
	ChangeTrainTypeName string `json:"change_train_type_name,omitempty" xml:"change_train_type_name,omitempty"`
	// 改签座位类型
	ChangeSeatTypeName string `json:"change_seat_type_name,omitempty" xml:"change_seat_type_name,omitempty"`
	// 改签票启程时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 改签票到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 进站时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 出站时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 外部票状态
	OutTicketStatus string `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	// 改签手续费
	ChangeHandlingFee float64 `json:"change_handling_fee,omitempty" xml:"change_handling_fee,omitempty"`
	// 改签差价
	ChangeGapFee float64 `json:"change_gap_fee,omitempty" xml:"change_gap_fee,omitempty"`
	// 改签服务费
	ChangeServiceFee float64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
}
