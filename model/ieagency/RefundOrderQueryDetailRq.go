package ieagency

// RefundOrderQueryDetailRq 
type RefundOrderQueryDetailRq struct {
    // 代理商ID
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    // 申请单ID
    RefundOrderId   int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}
