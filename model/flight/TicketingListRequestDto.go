package flight

// TicketingListRequestDto 
type TicketingListRequestDto struct {

    // 店铺id
    AgentIds   []Number `json:"agent_ids,omitempty"`

    // 支付起始时间
    BeginPayTime   string `json:"begin_pay_time,omitempty"`

    // 支付结束时间
    EndPayTime   string `json:"end_pay_time,omitempty"`

    // 页码
    Page   int64 `json:"page,omitempty"`

    // 状态码
    Status   int64 `json:"status,omitempty"`

}
