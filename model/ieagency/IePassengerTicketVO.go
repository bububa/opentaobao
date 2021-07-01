package ieagency

// IePassengerTicketVO 结构体
type IePassengerTicketVO struct {
	// 乘机人姓名（英文）
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 证件号码
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 票号
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 代理人pnr
	Cpnr string `json:"cpnr,omitempty" xml:"cpnr,omitempty"`
	// 航空公司pnr
	Bpnr string `json:"bpnr,omitempty" xml:"bpnr,omitempty"`
	// 出票渠道，可以填1E、1A、1S、1G、官网、其他
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}
