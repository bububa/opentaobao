package jipiao

// ReturnApplyDO 
type ReturnApplyDO struct {
    // 到达时间
    DepTime   string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
    // 退款金额（单位：元）
    RefundMoney   int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
    // 舱位
    Cabin   string `json:"cabin,omitempty" xml:"cabin,omitempty"`
    // 退款成功时间
    PaySuccessTime   string `json:"pay_success_time,omitempty" xml:"pay_success_time,omitempty"`
    // 退票手续费（单位：元）
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 退票提交时间
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    // 出发机场三字码
    DepAirportCode   string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
    // 数据项id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 票号
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    // 航线二字码
    AirlineCode   string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
    // 航班号
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
    // 退票成功时间
    FirstProcessTime   string `json:"first_process_time,omitempty" xml:"first_process_time,omitempty"`
    // 乘机人姓名
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 到达机场三字码
    ArrAirportCode   string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 退票状态，1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
