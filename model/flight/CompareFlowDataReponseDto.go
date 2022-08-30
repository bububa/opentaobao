package flight

// CompareFlowDataReponseDto 结构体
type CompareFlowDataReponseDto struct {
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 当日已使用量
	CurrentAmount int64 `json:"current_amount,omitempty" xml:"current_amount,omitempty"`
	// 当日分配给商家的总调用量
	TotalLimit int64 `json:"total_limit,omitempty" xml:"total_limit,omitempty"`
}
