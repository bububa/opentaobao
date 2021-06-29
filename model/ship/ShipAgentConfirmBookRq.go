package ship

// ShipAgentConfirmBookRq 
type ShipAgentConfirmBookRq struct {
    // 商家订单id
    AgentOrderId   string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
    // 飞猪订单id
    AlitripOrderId   string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
    // 取票地址信息
    FetchTicketsAddress   string `json:"fetch_tickets_address,omitempty" xml:"fetch_tickets_address,omitempty"`
    // 取票号
    FetchTicketsNumber   string `json:"fetch_tickets_number,omitempty" xml:"fetch_tickets_number,omitempty"`
    // 取票密码
    FetchTicketsPwd   string `json:"fetch_tickets_pwd,omitempty" xml:"fetch_tickets_pwd,omitempty"`
    // 淘宝订单Id
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
    // 取票短信内容
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 扩展属性
    OrderAttr   string `json:"order_attr,omitempty" xml:"order_attr,omitempty"`
    // 乘客列表
    PassengerList   []ShipAgentConfirmBookPassengerInfo `json:"passenger_list,omitempty" xml:"passenger_list>ship_agent_confirm_book_passenger_info,omitempty"`
    // 出票结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 票数
    TicketCount   int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
    // 检票口
    TicketWicket   string `json:"ticket_wicket,omitempty" xml:"ticket_wicket,omitempty"`
    // 总价
    TotalPrice   int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 出票失败错误码
    FailedCode   string `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
}
