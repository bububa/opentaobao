package bus

// B2BBusOrderInfo 
/* model for simplify = false
type B2BBusOrderInfo struct {

    // 支付宝交易流水号
    
    AlipayTradeId   string `json:"alipay_trade_id,omitempty"`
    

    // 车次对象
    
    B2BBusLineInfo  *struct {
        B2BBusLineInfo  *B2BBusLineInfo `json:"b2b_bus_line_info,omitempty"`
    } `json:"b2_b_bus_line_info,omitempty"`
    

    // 取票人信息
    
    B2BFetchHolderInfo  *struct {
        B2BFetchHolderInfo  *B2BFetchHolderInfo `json:"b2b_fetch_holder_info,omitempty"`
    } `json:"b2_b_fetch_holder_info,omitempty"`
    

    // 取票信息
    
    B2BFetchTicket  *struct {
        B2BFetchTicket  *B2BFetchTicket `json:"b2b_fetch_ticket,omitempty"`
    } `json:"b2_b_fetch_ticket,omitempty"`
    

    // 票信息
    
    BusB2bTicketInfoList  struct {
        B2BTicketInfo  []B2BTicketInfo `json:"b2b_ticket_info,omitempty"`
    } `json:"bus_b2b_ticket_info_list,omitempty"`
    

    // 订单创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 出票成功时间
    
    IssueTime   string `json:"issue_time,omitempty"`
    

    // 阿里订单号
    
    MainOrderId   int64 `json:"main_order_id,omitempty"`
    

    // 订单状态
    
    OrderStatus   int64 `json:"order_status,omitempty"`
    

    // 票数
    
    TicketCount   int64 `json:"ticket_count,omitempty"`
    

    // 总价
    
    TotalPrice   int64 `json:"total_price,omitempty"`
    

}
*/

// B2BBusOrderInfo 
type B2BBusOrderInfo struct {

    // 支付宝交易流水号
    AlipayTradeId   string `json:"alipay_trade_id,omitempty"`

    // 车次对象
    B2BBusLineInfo   *B2BBusLineInfo `json:"b2_b_bus_line_info,omitempty"`

    // 取票人信息
    B2BFetchHolderInfo   *B2BFetchHolderInfo `json:"b2_b_fetch_holder_info,omitempty"`

    // 取票信息
    B2BFetchTicket   *B2BFetchTicket `json:"b2_b_fetch_ticket,omitempty"`

    // 票信息
    BusB2bTicketInfoList   []B2BTicketInfo `json:"bus_b2b_ticket_info_list,omitempty"`

    // 订单创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 出票成功时间
    IssueTime   string `json:"issue_time,omitempty"`

    // 阿里订单号
    MainOrderId   int64 `json:"main_order_id,omitempty"`

    // 订单状态
    OrderStatus   int64 `json:"order_status,omitempty"`

    // 票数
    TicketCount   int64 `json:"ticket_count,omitempty"`

    // 总价
    TotalPrice   int64 `json:"total_price,omitempty"`

}
