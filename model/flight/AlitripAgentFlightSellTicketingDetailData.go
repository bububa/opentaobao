package flight

// AlitripAgentFlightSellTicketingDetailData 
/* model for simplify = false
type AlitripAgentFlightSellTicketingDetailData struct {

    // 店铺id
    
    AgentId   int64 `json:"agent_id,omitempty"`
    

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 飞猪订单号
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 出票时间
    
    IssueTime   string `json:"issue_time,omitempty"`
    

    // 出票对象
    
    IssueList  struct {
        IssueList  []IssueList `json:"issue_list,omitempty"`
    } `json:"issue_list,omitempty"`
    

    // sla
    
    Sla   string `json:"sla,omitempty"`
    

    // 佣金
    
    Commission   int64 `json:"commission,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty"`
    

    // 订单标签
    
    Tags  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tags,omitempty"`
    

    // 订单状态
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// AlitripAgentFlightSellTicketingDetailData 
type AlitripAgentFlightSellTicketingDetailData struct {

    // 店铺id
    AgentId   int64 `json:"agent_id,omitempty"`

    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 飞猪订单号
    OrderId   string `json:"order_id,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 出票时间
    IssueTime   string `json:"issue_time,omitempty"`

    // 出票对象
    IssueList   []IssueList `json:"issue_list,omitempty"`

    // sla
    Sla   string `json:"sla,omitempty"`

    // 佣金
    Commission   int64 `json:"commission,omitempty"`

    // 币种
    Currency   string `json:"currency,omitempty"`

    // 订单标签
    Tags   []string `json:"tags,omitempty"`

    // 订单状态
    Status   int64 `json:"status,omitempty"`

}
