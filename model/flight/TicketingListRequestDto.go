package flight

// TicketingListRequestDto 
/* model for simplify = false
type TicketingListRequestDto struct {

    // 店铺id
    
    AgentIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"agent_ids,omitempty"`
    

    // 支付起始时间
    
    BeginPayTime   string `json:"begin_pay_time,omitempty"`
    

    // 支付结束时间
    
    EndPayTime   string `json:"end_pay_time,omitempty"`
    

    // 页码
    
    Page   int64 `json:"page,omitempty"`
    

    // 状态码
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// TicketingListRequestDto 
type TicketingListRequestDto struct {

    // 店铺id
    AgentIds   []int64 `json:"agent_ids,omitempty"`

    // 支付起始时间
    BeginPayTime   string `json:"begin_pay_time,omitempty"`

    // 支付结束时间
    EndPayTime   string `json:"end_pay_time,omitempty"`

    // 页码
    Page   int64 `json:"page,omitempty"`

    // 状态码
    Status   int64 `json:"status,omitempty"`

}
