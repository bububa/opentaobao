package btrip

// OrderTicketInfo 结构体
type OrderTicketInfo struct {
	// pnr编码
	PnrCode string `json:"pnr_code,omitempty" xml:"pnr_code,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 票状态
	TicketStatus string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// openTicket状态
	OpenTicketStatus string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
}
