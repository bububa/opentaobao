package ieagency

// IeIssueTicketVo 结构体
type IeIssueTicketVo struct {
	// 乘机人票信息
	PassengerTicketVos []IePassengerTicketVo `json:"passenger_ticket_vos,omitempty" xml:"passenger_ticket_vos>ie_passenger_ticket_vo,omitempty"`
	// 预定订单pnr信息
	UpdatePnrVos []IeBookPnrVo `json:"update_pnr_vos,omitempty" xml:"update_pnr_vos>ie_book_pnr_vo,omitempty"`
	// 订单备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 预定订单id
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
}
