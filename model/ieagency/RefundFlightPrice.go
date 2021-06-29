package ieagency

// RefundFlightPrice 
type RefundFlightPrice struct {

    // 税费价格(单位：分)
    
    TaxPrice   int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
    

    // 机票价格(单位:分)
    
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    

}
