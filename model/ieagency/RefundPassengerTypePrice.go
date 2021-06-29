package ieagency

// RefundPassengerTypePrice 
type RefundPassengerTypePrice struct {
    // 机票价格信息
    FlightPrice   *RefundFlightPrice `json:"flight_price,omitempty" xml:"flight_price,omitempty"`
    // 乘机人类型(Adult(0, "成人"),     Child(1, "儿童"),     StudentAbroad(2, "留学生"),     Infant(3, "婴儿")
    PassengerType   int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
