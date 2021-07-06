package ieagency

// IeChangeTicketVo 结构体
type IeChangeTicketVo struct {
	// 证件号
	DocumentID string `json:"document_i_d,omitempty" xml:"document_i_d,omitempty"`
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 票号
	TicketNO string `json:"ticket_n_o,omitempty" xml:"ticket_n_o,omitempty"`
	// 乘机人类型  0-成人,1-儿童,2-留学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
