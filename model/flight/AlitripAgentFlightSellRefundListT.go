package flight

// AlitripAgentFlightSellRefundListT 
type AlitripAgentFlightSellRefundListT struct {

    // 退票申请单号
    
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
    

    // 飞猪订单号
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 时间
    
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    

}
