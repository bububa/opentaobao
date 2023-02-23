package bus

// MerchantBusTicketInfo 结构体
type MerchantBusTicketInfo struct {
	// 商家票id
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 座位号
	RiderSeatNumber string `json:"rider_seat_number,omitempty" xml:"rider_seat_number,omitempty"`
	// 乘客姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 商家服务费
	ServiceCharge int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
	// 子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 票价，单位分
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 0-全票(成人) 1-半票（儿童） 2-免票（携童）
	TicketType int64 `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
}
