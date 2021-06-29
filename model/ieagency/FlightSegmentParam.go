package ieagency

// FlightSegmentParam 
type FlightSegmentParam struct {

    // 到达机场三字码
    
    ArrAirportCode   string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
    

    // 到达机场航站楼
    
    ArrTerminal   string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
    

    // 到达时间
    
    ArrTime   string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
    

    // 舱位服务等级
    
    CabinClassCode   string `json:"cabin_class_code,omitempty" xml:"cabin_class_code,omitempty"`
    

    // 舱位等级
    
    CabinCode   string `json:"cabin_code,omitempty" xml:"cabin_code,omitempty"`
    

    // 出发机场三字码
    
    DepAirportCode   string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
    

    // 出发机场航站楼
    
    DepTerminal   string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
    

    // 出发时间
    
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    

    // 飞行时长(分钟)
    
    ElapsedMinute   int64 `json:"elapsed_minute,omitempty" xml:"elapsed_minute,omitempty"`
    

    // 机型
    
    EquipType   string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
    

    // 市场方航空公司
    
    MarketingAirline   string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
    

    // 市场方航班号
    
    MarketingFlightNumber   string `json:"marketing_flight_number,omitempty" xml:"marketing_flight_number,omitempty"`
    

    // 承运航空公司
    
    OperatingAirLine   string `json:"operating_air_line,omitempty" xml:"operating_air_line,omitempty"`
    

    // 承运航班号
    
    OperatingFlightNumber   string `json:"operating_flight_number,omitempty" xml:"operating_flight_number,omitempty"`
    

    // 航段序号(从1开始)
    
    SegmentRph   int64 `json:"segment_rph,omitempty" xml:"segment_rph,omitempty"`
    

    // 主航段
    
    MainSegment   bool `json:"main_segment,omitempty" xml:"main_segment,omitempty"`
    

}
