package flight

// RefundListRequestDto 
type RefundListRequestDto struct {

    // 退票申请起始时间
    EndApplyTime   string `json:"end_apply_time,omitempty"`

    // 退票申请结束时间
    BeginApplyTime   string `json:"begin_apply_time,omitempty"`

    // 页码
    Page   int64 `json:"page,omitempty"`

    // 退票单状态
    Status   int64 `json:"status,omitempty"`

    // 店铺id集合
    AgentIds   []Number `json:"agent_ids,omitempty"`

}
