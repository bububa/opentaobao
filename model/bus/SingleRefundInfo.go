package bus

// SingleRefundInfo 
/* model for simplify = false
type SingleRefundInfo struct {

    // 乘客身份证后四位
    
    CardNo   string `json:"card_no,omitempty"`
    

    // 退票价格
    
    RefundPrice   string `json:"refund_price,omitempty"`
    

    // 乘客姓名
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 退票手续费
    
    RefundProcedurePrice   string `json:"refund_procedure_price,omitempty"`
    

    // 退款详情
    
    RefundDetail   string `json:"refund_detail,omitempty"`
    

    // 代理商订单号
    
    AgentOrderId   string `json:"agent_order_id,omitempty"`
    

    // 退票时间
    
    RefundTicketTime   string `json:"refund_ticket_time,omitempty"`
    

    // 代理商票ID
    
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`
    

}
*/

// SingleRefundInfo 
type SingleRefundInfo struct {

    // 乘客身份证后四位
    CardNo   string `json:"card_no,omitempty"`

    // 退票价格
    RefundPrice   string `json:"refund_price,omitempty"`

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 退票手续费
    RefundProcedurePrice   string `json:"refund_procedure_price,omitempty"`

    // 退款详情
    RefundDetail   string `json:"refund_detail,omitempty"`

    // 代理商订单号
    AgentOrderId   string `json:"agent_order_id,omitempty"`

    // 退票时间
    RefundTicketTime   string `json:"refund_ticket_time,omitempty"`

    // 代理商票ID
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`

}
