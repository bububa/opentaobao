package btrip

// OpenTicketDo 
type OpenTicketDo struct {
    // 机票号
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    // 是否改签票（true：是；false：否）
    IsChanged   bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
    // 原改签票号
    OriginTicketNo   string `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
    // 航班号
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
    // 航空公司名称
    AirlineCompany   string `json:"airline_company,omitempty" xml:"airline_company,omitempty"`
    // 航空公司编码
    AirlineCode   string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
    // 出发时间
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    // 出发城市
    DepCityName   string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
    // 到达城市
    ArrCityName   string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
    // 出发机场三字码
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    // 到达机场三字码
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    // 出发机场名称
    DepAirportName   string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
    // 到达机场名称
    ArrAirportName   string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
    // 舱位
    Cabin   string `json:"cabin,omitempty" xml:"cabin,omitempty"`
    // 舱等
    Cabinclass   string `json:"cabinclass,omitempty" xml:"cabinclass,omitempty"`
    // 票价（分）
    TicketPrice   string `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    // 折扣
    Discount   string `json:"discount,omitempty" xml:"discount,omitempty"`
    // 保险费（分）
    InsuranceFee   int64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
    // 燃油费（分）
    Oil   int64 `json:"oil,omitempty" xml:"oil,omitempty"`
    // 基建（分）
    Build   int64 `json:"build,omitempty" xml:"build,omitempty"`
    // 货币类型
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 行程单号
    ItineraryNum   string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
    // 乘机人
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 保险单号
    InsureNo   string `json:"insure_no,omitempty" xml:"insure_no,omitempty"`
    // 状态：1已出保 2已退保
    InsureStatus   string `json:"insure_status,omitempty" xml:"insure_status,omitempty"`
    // 乘机人(保险人)姓名
    InsureName   string `json:"insure_name,omitempty" xml:"insure_name,omitempty"`
    // 飞行时长（分钟），未去除经停时间
    RideTime   int64 `json:"ride_time,omitempty" xml:"ride_time,omitempty"`
    // 到达时间
    ArrTime   string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
}
