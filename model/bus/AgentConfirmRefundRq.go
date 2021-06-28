package bus

// AgentConfirmRefundRq 
/* model for simplify = false
type AgentConfirmRefundRq struct {

    // 商家单号
    
    AgentOrderId   string `json:"agent_order_id,omitempty"`
    

    // 按票退时商家退款票号
    
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`
    

    // 发车时间
    
    DepartDate   string `json:"depart_date,omitempty"`
    

    // 平台单号
    
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty"`
    

    // 退款乘客身份证号
    
    PassengerIdNum   string `json:"passenger_id_num,omitempty"`
    

    // 退款乘客手机号
    
    PassengerPhone   string `json:"passenger_phone,omitempty"`
    

    // 退款金额，单位（元）
    
    RefundFee   string `json:"refund_fee,omitempty"`
    

    // 退款时间点
    
    RefundTime   string `json:"refund_time,omitempty"`
    

    // 退款资金号唯一ID
    
    RefundTransId   string `json:"refund_trans_id,omitempty"`
    

    // 退款类型 0-退票
    
    RefundType   int64 `json:"refund_type,omitempty"`
    

}
*/

// AgentConfirmRefundRq 
type AgentConfirmRefundRq struct {

    // 商家单号
    AgentOrderId   string `json:"agent_order_id,omitempty"`

    // 按票退时商家退款票号
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`

    // 发车时间
    DepartDate   string `json:"depart_date,omitempty"`

    // 平台单号
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty"`

    // 退款乘客身份证号
    PassengerIdNum   string `json:"passenger_id_num,omitempty"`

    // 退款乘客手机号
    PassengerPhone   string `json:"passenger_phone,omitempty"`

    // 退款金额，单位（元）
    RefundFee   string `json:"refund_fee,omitempty"`

    // 退款时间点
    RefundTime   string `json:"refund_time,omitempty"`

    // 退款资金号唯一ID
    RefundTransId   string `json:"refund_trans_id,omitempty"`

    // 退款类型 0-退票
    RefundType   int64 `json:"refund_type,omitempty"`

}
