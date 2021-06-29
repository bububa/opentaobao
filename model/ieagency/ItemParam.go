package ieagency

// ItemParam 
type ItemParam struct {

    // 成人税费（单位：分）
    
    AdultTax   int64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
    

    // 成人票价（单位：分）
    
    AdultTicketPrice   int64 `json:"adult_ticket_price,omitempty" xml:"adult_ticket_price,omitempty"`
    

    // 儿童税费（单位：分）
    
    ChildTax   int64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
    

    // 儿童票价（单位：分）
    
    ChildTicketPrice   int64 `json:"child_ticket_price,omitempty" xml:"child_ticket_price,omitempty"`
    

    // 航班行程信息
    
    ItineraryParams   []ItineraryParam `json:"itinerary_params,omitempty" xml:"itinerary_params,omitempty"`
    

    // 行程类型(1:单程;2:往返; 3:多程）
    
    TripType   int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
    

}
