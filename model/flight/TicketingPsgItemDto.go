package flight

// TicketingPsgItemDto 
type TicketingPsgItemDto struct {
    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 票号
    Tickets   []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
    // pnr
    Pnr   string `json:"pnr,omitempty" xml:"pnr,omitempty"`
    // 航段
    Segments   []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
}
