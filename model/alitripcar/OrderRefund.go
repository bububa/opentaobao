package alitripcar

// OrderRefund 
type OrderRefund struct {

    // 退款金额(单位：分)
    
    RefundFee   string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    

    // 订单原始金额(单位：分)
    
    OriginalPrice   string `json:"original_price,omitempty" xml:"original_price,omitempty"`
    

    // 服务商订单id
    
    ThirdOrderId   string `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
    

    // 飞猪订单id
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 商家退款唯一标识(密等对账使用)
    
    AgentUniqKey   string `json:"agent_uniq_key,omitempty" xml:"agent_uniq_key,omitempty"`
    

    // 供应商编号
    
    ProviderId   string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
    

}
