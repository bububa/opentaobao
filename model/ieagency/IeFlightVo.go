package ieagency

// IeFlightVo 
type IeFlightVo struct {

    // 到达机场
    
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    

    // 服务等级(First:头等舱,Business:商务舱, Economy:经济舱, EconomyStandard:标准经济舱, EconomyPremium:超级经济舱)
    
    CabinClass   string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
    

    // 出发时间
    
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    

    // 经停数量
    
    StopQuantity   int64 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
    

    // 舱位
    
    FlightCabin   string `json:"flight_cabin,omitempty" xml:"flight_cabin,omitempty"`
    

    // 过境签信息
    
    TransVisa   string `json:"trans_visa,omitempty" xml:"trans_visa,omitempty"`
    

    // 出发机场
    
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    

    // 出发航站楼
    
    DepTerminal   string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
    

    // 航班号
    
    FlightNumber   string `json:"flight_number,omitempty" xml:"flight_number,omitempty"`
    

    // 机型
    
    EquipType   string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
    

    // 经停机场三字码，逗号分隔
    
    StopAirport   string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
    

    // 第几段
    
    SegmentRph   int64 `json:"segment_rph,omitempty" xml:"segment_rph,omitempty"`
    

    // 实际承运航班号
    
    OperatingFlightNumber   string `json:"operating_flight_number,omitempty" xml:"operating_flight_number,omitempty"`
    

    // 承运方航空公司
    
    OperatingAirLine   string `json:"operating_air_line,omitempty" xml:"operating_air_line,omitempty"`
    

    // 飞行时长
    
    ElapsedMinute   int64 `json:"elapsed_minute,omitempty" xml:"elapsed_minute,omitempty"`
    

    // 到达航站楼
    
    ArrTerminal   string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
    

    // 到达时间
    
    ArrTime   string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
    

    // 经停分钟数，逗号分隔
    
    StopMinute   int64 `json:"stop_minute,omitempty" xml:"stop_minute,omitempty"`
    

    // 行程类型
    
    ItineraryType   int64 `json:"itinerary_type,omitempty" xml:"itinerary_type,omitempty"`
    

    // 市场方航空公司
    
    MarketingAirline   string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
    

    // cabinClassCode
    
    CabinClassCode   string `json:"cabin_class_code,omitempty" xml:"cabin_class_code,omitempty"`
    

    // mainSegment
    
    MainSegment   bool `json:"main_segment,omitempty" xml:"main_segment,omitempty"`
    

    // 婴儿舱位等级代码
    
    InfantCabinClassCode   string `json:"infant_cabin_class_code,omitempty" xml:"infant_cabin_class_code,omitempty"`
    

    // 服务等级(First:头等舱,Business:商务舱, Economy:经济舱, EconomyStandard:标准经济舱, EconomyPremium:超级经济舱)
    
    InfantCabinClass   string `json:"infant_cabin_class,omitempty" xml:"infant_cabin_class,omitempty"`
    

    // 婴儿舱位代码
    
    InfantFlightCabin   string `json:"infant_flight_cabin,omitempty" xml:"infant_flight_cabin,omitempty"`
    

}
