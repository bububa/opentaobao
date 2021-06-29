package ieagency

// RefundItemVo 
type RefundItemVo struct {
    // 乘机人类型价格信息
    PassengerTypePrices   []RefundPassengerTypePrice `json:"passenger_type_prices,omitempty" xml:"passenger_type_prices>refund_passenger_type_price,omitempty"`
    // 行程信息
    RefundItineraryFlights   []RefundItineraryVo `json:"refund_itinerary_flights,omitempty" xml:"refund_itinerary_flights>refund_itinerary_vo,omitempty"`
}
