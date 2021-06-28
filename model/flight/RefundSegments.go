package flight

// RefundSegments 
/* model for simplify = false
type RefundSegments struct {

    // 舱等
    
    CabinClass   string `json:"cabin_class,omitempty"`
    

    // 航班号
    
    FlightNo   string `json:"flight_no,omitempty"`
    

    // 起飞时间
    
    DepTime   string `json:"dep_time,omitempty"`
    

    // 到达城市
    
    ArrCity   string `json:"arr_city,omitempty"`
    

    // 起飞城市
    
    DepCity   string `json:"dep_city,omitempty"`
    

    // 舱位
    
    Cabin   string `json:"cabin,omitempty"`
    

    // 航段序号
    
    SegmentIndex   int64 `json:"segment_index,omitempty"`
    

    // 到达机场
    
    ArrAirport   string `json:"arr_airport,omitempty"`
    

    // 起飞机场
    
    DepAirport   string `json:"dep_airport,omitempty"`
    

    // 到达时间
    
    ArrTime   string `json:"arr_time,omitempty"`
    

    // 航程序号
    
    OdIndex   int64 `json:"od_index,omitempty"`
    

}
*/

// RefundSegments 
type RefundSegments struct {

    // 舱等
    CabinClass   string `json:"cabin_class,omitempty"`

    // 航班号
    FlightNo   string `json:"flight_no,omitempty"`

    // 起飞时间
    DepTime   string `json:"dep_time,omitempty"`

    // 到达城市
    ArrCity   string `json:"arr_city,omitempty"`

    // 起飞城市
    DepCity   string `json:"dep_city,omitempty"`

    // 舱位
    Cabin   string `json:"cabin,omitempty"`

    // 航段序号
    SegmentIndex   int64 `json:"segment_index,omitempty"`

    // 到达机场
    ArrAirport   string `json:"arr_airport,omitempty"`

    // 起飞机场
    DepAirport   string `json:"dep_airport,omitempty"`

    // 到达时间
    ArrTime   string `json:"arr_time,omitempty"`

    // 航程序号
    OdIndex   int64 `json:"od_index,omitempty"`

}
