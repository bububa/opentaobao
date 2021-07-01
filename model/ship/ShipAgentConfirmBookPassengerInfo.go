package ship

// ShipAgentConfirmBookPassengerInfo 结构体
type ShipAgentConfirmBookPassengerInfo struct {
	// 乘客证件号
	PassengerCertNo string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// 乘客证件类型
	PassengerCertType string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// 乘客id
	PassengerId string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 票信息
	TicketList []ShipAgentConfirmBookTicketInfo `json:"ticket_list,omitempty" xml:"ticket_list>ship_agent_confirm_book_ticket_info,omitempty"`
}
