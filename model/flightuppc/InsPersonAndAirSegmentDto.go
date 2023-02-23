package flightuppc

// InsPersonAndAirSegmentDto 结构体
type InsPersonAndAirSegmentDto struct {
	// 航段信息
	InsOrderSegment *InsOrderAirTicketSegmentDto `json:"ins_order_segment,omitempty" xml:"ins_order_segment,omitempty"`
	// 被保人
	InsPerson *InsPersonDto `json:"ins_person,omitempty" xml:"ins_person,omitempty"`
}
