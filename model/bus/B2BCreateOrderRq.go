package bus

// B2bcreateOrderRq 结构体
type B2bcreateOrderRq struct {
	// 乘客信息
	Passengers []PassengerVo `json:"passengers,omitempty" xml:"passengers>passenger_vo,omitempty"`
	// 取票人
	B2BFetchHolderInfo *B2bfetchHolderInfo `json:"b2_b_fetch_holder_info,omitempty" xml:"b2_b_fetch_holder_info,omitempty"`
	// 车次信息
	B2bBusLineInfo *B2bbusLineInfo `json:"b2b_bus_line_info,omitempty" xml:"b2b_bus_line_info,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}
