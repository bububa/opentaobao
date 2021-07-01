package flight

// ModifyItemDto 结构体
type ModifyItemDto struct {
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 改签后航段
	AfterChangeSegments []ModifySegmentDto `json:"after_change_segments,omitempty" xml:"after_change_segments>modify_segment_dto,omitempty"`
	// 改签费用
	ModifyFee int64 `json:"modify_fee,omitempty" xml:"modify_fee,omitempty"`
	// 升舱费用
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// 改前航段
	BeforeChangeSegments []ModifyBeforeSegmentDto `json:"before_change_segments,omitempty" xml:"before_change_segments>modify_before_segment_dto,omitempty"`
	// 票号
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
}
