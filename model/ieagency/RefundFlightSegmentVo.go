package ieagency

// RefundFlightSegmentVo 
type RefundFlightSegmentVo struct {

    // 到达机场
    
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    

    // 到达城市
    
    ArrCity   string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
    

    // 到达航站楼
    
    ArrTerminal   string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
    

    // 到达时间
    
    ArrTime   string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
    

    // 共享编码
    
    CodeShare   bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
    

    // 出发机场
    
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    

    // 出发城市
    
    DepCity   string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
    

    // 出发航站楼
    
    DepTerminal   string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
    

    // 出发时间
    
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    

    // 航段ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 行程序号
    
    ItineraryIndex   int64 `json:"itinerary_index,omitempty" xml:"itinerary_index,omitempty"`
    

    // 主航段
    
    MainSegment   bool `json:"main_segment,omitempty" xml:"main_segment,omitempty"`
    

    // 市场方航司
    
    MarketingAirline   string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
    

    // 市场方航班号
    
    MarketingFlightNumber   string `json:"marketing_flight_number,omitempty" xml:"marketing_flight_number,omitempty"`
    

    // 执行航班号
    
    OperatingFlightNumber   string `json:"operating_flight_number,omitempty" xml:"operating_flight_number,omitempty"`
    

    // 航段序号
    
    SegmentIndex   int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
    

    // 执行方航班号
    
    OperatingAirline   string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
    

}
