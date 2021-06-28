package flight

// TicketingListRequestDto 
type TicketingListRequestDto struct {

    // 店铺id
    
    AgentIds   []int64 `json:"agent_ids,omitempty" xml:"agent_ids>int64,omitempty"`
    

    // 支付起始时间
    
    BeginPayTime   string `json:"begin_pay_time,omitempty" xml:"begin_pay_time,omitempty"`
    

    // 支付结束时间
    
    EndPayTime   string `json:"end_pay_time,omitempty" xml:"end_pay_time,omitempty"`
    

    // 页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 状态码
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}
