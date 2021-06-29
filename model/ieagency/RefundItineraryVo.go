package ieagency

// RefundItineraryVo 
type RefundItineraryVo struct {

    // 出发机场
    
    DepAirportCode   string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
    

    // 出发时间
    
    DepDate   string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
    

    // 行程序号
    
    Index   int64 `json:"index,omitempty" xml:"index,omitempty"`
    

    // 航班列表
    
    RefundFlightSegmentVos   []RefundFlightSegmentVo `json:"refund_flight_segment_vos,omitempty" xml:"refund_flight_segment_vos,omitempty"`
    

    // 到达机场
    
    ArrAirportCode   string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
    

}
