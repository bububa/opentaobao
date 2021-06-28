package flight

// TicketingPsgItemDto 
/* model for simplify = false
type TicketingPsgItemDto struct {

    // 乘客姓名
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 票号
    
    Tickets  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tickets,omitempty"`
    

    // pnr
    
    Pnr   string `json:"pnr,omitempty"`
    

    // 航段
    
    Segments  struct {
        Segments  []Segments `json:"segments,omitempty"`
    } `json:"segments,omitempty"`
    

}
*/

// TicketingPsgItemDto 
type TicketingPsgItemDto struct {

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 票号
    Tickets   []string `json:"tickets,omitempty"`

    // pnr
    Pnr   string `json:"pnr,omitempty"`

    // 航段
    Segments   []Segments `json:"segments,omitempty"`

}
