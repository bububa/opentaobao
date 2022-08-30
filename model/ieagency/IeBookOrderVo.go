package ieagency

// IeBookOrderVo 结构体
type IeBookOrderVo struct {
	// 2
	BookFlightVos []IeBookFlightVo `json:"book_flight_vos,omitempty" xml:"book_flight_vos>ie_book_flight_vo,omitempty"`
	// 9
	BookPnrVos []IeBookPnrVo `json:"book_pnr_vos,omitempty" xml:"book_pnr_vos>ie_book_pnr_vo,omitempty"`
	// 13
	BookTicketVos []IeBookTicketVo `json:"book_ticket_vos,omitempty" xml:"book_ticket_vos>ie_book_ticket_vo,omitempty"`
	// 预定订单状态(Init:初始状态, Need_Manule_HK:转手工HK, HK_Success:HK成功,Ticket_Success:出票成功,Close:订单关闭)
	BookStatus string `json:"book_status,omitempty" xml:"book_status,omitempty"`
	// pnr预定类型(Auto:自动, Manual:手工)
	BookType string `json:"book_type,omitempty" xml:"book_type,omitempty"`
	// 出票类型(Auto:自动出票,Manual:手工出票)
	IssueTicketType string `json:"issue_ticket_type,omitempty" xml:"issue_ticket_type,omitempty"`
	// 税费类型(UserRefTax:参考税,RealTax:实时税费,ManualTax:手工回填税费)
	TaxType string `json:"tax_type,omitempty" xml:"tax_type,omitempty"`
	// 19
	TerminalFileName string `json:"terminal_file_name,omitempty" xml:"terminal_file_name,omitempty"`
	// 18
	TerminalUrl string `json:"terminal_url,omitempty" xml:"terminal_url,omitempty"`
	// 票号验真状态(Init:初始状态,Failure:验真失败,SUCCESS:验真通过)
	TicketAuthenticStatus string `json:"ticket_authentic_status,omitempty" xml:"ticket_authentic_status,omitempty"`
	// 15
	AdtTaxPrice int64 `json:"adt_tax_price,omitempty" xml:"adt_tax_price,omitempty"`
	// 10
	AdtTicketPrice int64 `json:"adt_ticket_price,omitempty" xml:"adt_ticket_price,omitempty"`
	// 300
	ChdTaxPrice int64 `json:"chd_tax_price,omitempty" xml:"chd_tax_price,omitempty"`
	// 1
	ChdTicketPrice int64 `json:"chd_ticket_price,omitempty" xml:"chd_ticket_price,omitempty"`
	// 3
	FareSource int64 `json:"fare_source,omitempty" xml:"fare_source,omitempty"`
	// 原始总价(包含打包商品的价格)
	OriginTotalPrice int64 `json:"origin_total_price,omitempty" xml:"origin_total_price,omitempty"`
	// 16
	PassengerCount int64 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	// 5000
	TotalOrderPirce int64 `json:"total_order_pirce,omitempty" xml:"total_order_pirce,omitempty"`
	// 17
	TotalTaxPrice int64 `json:"total_tax_price,omitempty" xml:"total_tax_price,omitempty"`
	// 12
	TotalTicketPrice int64 `json:"total_ticket_price,omitempty" xml:"total_ticket_price,omitempty"`
}
