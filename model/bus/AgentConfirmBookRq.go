package bus

// AgentConfirmBookRq 
type AgentConfirmBookRq struct {

    // 代理商订单号
    AgentOrderId   string `json:"agent_order_id,omitempty"`

    // 座位号，多个座位号以","分隔
    RiderSeatNumbers   string `json:"rider_seat_numbers,omitempty"`

    // 短信内容。商家如有特殊定制，则需要严格按照参数返回json格式数据。{"driverContactPhone":"18611330989","identificationCode","988838389","ticketInstructions":"套票使用说明"} driverContactPhone:司机联系电话 identificationCode:乘车验证码 ticketInstructions:门票使用说明（仅供门票+车票 类型商品使用，非必填，默认为“凭身份证入园”）
    Message   string `json:"message,omitempty"`

    // 票数
    TicketCount   int64 `json:"ticket_count,omitempty"`

    // 取票密码
    FetchTicketsPwd   string `json:"fetch_tickets_pwd,omitempty"`

    // 取票地点
    FetchTicketsAddress   string `json:"fetch_tickets_address,omitempty"`

    // 取票号
    FetchTicketsNumber   string `json:"fetch_tickets_number,omitempty"`

    // 是否出票成功
    Success   bool `json:"success,omitempty"`

    // 总价格
    TotalPrice   int64 `json:"total_price,omitempty"`

    // 1223dsd32323
    BusInnerOrderId   string `json:"bus_inner_order_id,omitempty"`

    // 检票口
    TicketWicket   string `json:"ticket_wicket,omitempty"`

    // 乘客信息
    PassengerInfoList   []AgentConfirmBookPassengerInfo `json:"passenger_info_list,omitempty"`

}
