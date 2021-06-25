package train

// CompensateParam 
type CompensateParam struct {

    // 代理商id
    AgentId   int64 `json:"agent_id,omitempty"`

    // 订单id
    OrderId   int64 `json:"order_id,omitempty"`

    // 改签单id
    ApplyId   int64 `json:"apply_id,omitempty"`

}
