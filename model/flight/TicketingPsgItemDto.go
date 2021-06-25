package flight

// TicketingPsgItemDto 
type TicketingPsgItemDto struct {

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 票号
    Tickets   []String `json:"tickets,omitempty"`

    // pnr
    Pnr   string `json:"pnr,omitempty"`

    // 航段
    Segments   []Segments `json:"segments,omitempty"`

}
