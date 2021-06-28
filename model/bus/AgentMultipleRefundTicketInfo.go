package bus

// AgentMultipleRefundTicketInfo 
type AgentMultipleRefundTicketInfo struct {

    // 商家票号(唯一标识关联乘客)
    
    AgentTicketId   string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
    

    // 退款人证件号(无票号时此参数必填)
    
    PassengerIdNum   string `json:"passenger_id_num,omitempty" xml:"passenger_id_num,omitempty"`
    

    // 退票金额(分)
    
    RefundAmount   int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
    

    // 退服务费金额(分)
    
    ServiceChargeRefundAmount   int64 `json:"service_charge_refund_amount,omitempty" xml:"service_charge_refund_amount,omitempty"`
    

}
