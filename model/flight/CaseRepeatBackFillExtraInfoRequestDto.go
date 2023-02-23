package flight

// CaseRepeatBackFillExtraInfoRequestDto 结构体
type CaseRepeatBackFillExtraInfoRequestDto struct {
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 人商品Id
	PassengerItemId int64 `json:"passenger_item_id,omitempty" xml:"passenger_item_id,omitempty"`
}
