package flight

// BookFlightVo 
type BookFlightVo struct {

    // 到达机场三字码
    
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    

    // 出发机场三字码
    
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    

    // 出发时间
    
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    

    // 仓位
    
    FlightCabin   string `json:"flight_cabin,omitempty" xml:"flight_cabin,omitempty"`
    

    // 航班号
    
    FlightNumber   string `json:"flight_number,omitempty" xml:"flight_number,omitempty"`
    

}
