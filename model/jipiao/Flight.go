package jipiao

// Flight 
type Flight struct {
    // 改签后航空公司二字码
    AirLineCode   string `json:"air_line_code,omitempty" xml:"air_line_code,omitempty"`
    // 改签后到达机场三字码
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    // 改签后出发机场三字码
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    // 改签后出发时间
    DepDate   string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
    // 改签后航班号
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}
