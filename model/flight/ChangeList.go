package flight

// ChangeList 
type ChangeList struct {
    // 证件号
    CertNo   string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 证件类型
    CertType   int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
    // 乘客类型
    PassengerType   int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    // 票号
    Tickets   []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
    // 优惠
    Promotion   int64 `json:"promotion,omitempty" xml:"promotion,omitempty"`
    // 改签前数据
    BeforeChangeSegments   []BeforeChangeSegments `json:"before_change_segments,omitempty" xml:"before_change_segments>before_change_segments,omitempty"`
    // 改签后数据
    AfterChangeSegments   []AfterChangeSegments `json:"after_change_segments,omitempty" xml:"after_change_segments>after_change_segments,omitempty"`
    // 票面价
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    // 改签费
    ChangeFee   int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
    // 升舱费
    UpgradeFee   int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
    // 改签后pnr
    Pnr   string `json:"pnr,omitempty" xml:"pnr,omitempty"`
    // 改签前票号
    BeforeChangeTickets   []string `json:"before_change_tickets,omitempty" xml:"before_change_tickets>string,omitempty"`
}
