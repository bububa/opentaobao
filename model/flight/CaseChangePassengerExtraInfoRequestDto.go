package flight

// CaseChangePassengerExtraInfoRequestDto 结构体
type CaseChangePassengerExtraInfoRequestDto struct {
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 手工单人ID
	ManualPassengerId int64 `json:"manual_passenger_id,omitempty" xml:"manual_passenger_id,omitempty"`
}
