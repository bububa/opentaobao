package bus

// B2BCreateOrderRq 
type B2BCreateOrderRq struct {

    // 取票人
    B2BFetchHolderInfo   *B2BFetchHolderInfo `json:"b2_b_fetch_holder_info,omitempty"`

    // 车次信息
    B2bBusLineInfo   *B2BBusLineInfo `json:"b2b_bus_line_info,omitempty"`

    // 乘客信息
    Passengers   []PassengerVo `json:"passengers,omitempty"`

    // 票数
    TicketCount   int64 `json:"ticket_count,omitempty"`

    // 总价
    TotalPrice   int64 `json:"total_price,omitempty"`

}
