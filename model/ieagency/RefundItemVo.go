package ieagency

// RefundItemVo 
type RefundItemVo struct {

    // 乘机人类型价格信息
    
    PassengerTypePrices   []RefundPassengerTypePrice `json:"passenger_type_prices,omitempty" xml:"passenger_type_prices,omitempty"`
    

    // 行程信息
    
    RefundItineraryFlights   []RefundItineraryVo `json:"refund_itinerary_flights,omitempty" xml:"refund_itinerary_flights,omitempty"`
    

}
