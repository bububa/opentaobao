package alihealth2

// StatementDetailVo 
type StatementDetailVo struct {

    // 订单ID
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 结算价，单位分
    
    SettlementPrice   string `json:"settlement_price,omitempty" xml:"settlement_price,omitempty"`
    

    // 成本，单位分
    
    CostPrice   string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
    

    // 利润，单位分
    
    ProfitPrice   string `json:"profit_price,omitempty" xml:"profit_price,omitempty"`
    

    // 核销状态 0 未核销 1已核销
    
    ConsumeStatus   int64 `json:"consume_status,omitempty" xml:"consume_status,omitempty"`
    

}
