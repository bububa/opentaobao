package jst

// OcOrder 
type OcOrder struct {

    // 分账总金额
    
    AllocatePaymentAmount   int64 `json:"allocate_payment_amount,omitempty" xml:"allocate_payment_amount,omitempty"`
    

    // 源订单号
    
    SourceTradeNo   string `json:"source_trade_no,omitempty" xml:"source_trade_no,omitempty"`
    

    // 淘宝账户id
    
    MainAllocationPaymentAcount   int64 `json:"main_allocation_payment_acount,omitempty" xml:"main_allocation_payment_acount,omitempty"`
    

    // 子订单集合
    
    OrderLines   []ChildOcOrder `json:"order_lines,omitempty" xml:"order_lines,omitempty"`
    

    // 来源渠道
    
    SourceTradeChannel   string `json:"source_trade_channel,omitempty" xml:"source_trade_channel,omitempty"`
    

    // 分账规则id
    
    AllocatePaymentRuleId   int64 `json:"allocate_payment_rule_id,omitempty" xml:"allocate_payment_rule_id,omitempty"`
    

    // 表示分账参与方
    
    AllocatePaymentParticipants   string `json:"allocate_payment_participants,omitempty" xml:"allocate_payment_participants,omitempty"`
    

    // OC订单号
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

}
