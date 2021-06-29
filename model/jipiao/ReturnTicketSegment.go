package jipiao

// ReturnTicketSegment 
type ReturnTicketSegment struct {

    // 到达机场三字码
    
    ArrAirportCode   string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
    

    // 到达城市
    
    ArrCity   string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
    

    // 单个航段机场建设费用（分）
    
    BuildFee   int64 `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
    

    // 出发机场三字码
    
    DepAirportCode   string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
    

    // 出发城市
    
    DepCity   string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
    

    // 航班号
    
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
    

    // 航段ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 单个航段机场建燃油费用
    
    OilTax   int64 `json:"oil_tax,omitempty" xml:"oil_tax,omitempty"`
    

    // 改签手续费(分)
    
    RefundModifyFee   int64 `json:"refund_modify_fee,omitempty" xml:"refund_modify_fee,omitempty"`
    

    // 升舱手续费（分）
    
    RefundUpgradeFee   int64 `json:"refund_upgrade_fee,omitempty" xml:"refund_upgrade_fee,omitempty"`
    

    // 票状态是否挂起
    
    Suspend   bool `json:"suspend,omitempty" xml:"suspend,omitempty"`
    

    // 票号信息
    
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    

    // 去程0 回程1
    
    TripType   int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
    

    // 起飞时间
    
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    

}
