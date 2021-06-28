package flight

// AlitripAgentFlightSellTicketingDetailData 
type AlitripAgentFlightSellTicketingDetailData struct {

    // 店铺id
    
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
    

    // 飞猪订单号
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 出票时间
    
    IssueTime   string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
    

    // 出票对象
    
    IssueList   []IssueList `json:"issue_list,omitempty" xml:"issue_list,omitempty"`
    

    // sla
    
    Sla   string `json:"sla,omitempty" xml:"sla,omitempty"`
    

    // 佣金
    
    Commission   int64 `json:"commission,omitempty" xml:"commission,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 订单标签
    
    Tags   []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
    

    // 订单状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}
