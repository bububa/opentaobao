package flight

// ConfirmRequestDto 结构体
type ConfirmRequestDto struct {
	// 出价，要么为null，要么集合内必须有值
	Prices []float64 `json:"prices,omitempty" xml:"prices>float64,omitempty"`
	// 商家店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 意向单id
	IntentionId int64 `json:"intention_id,omitempty" xml:"intention_id,omitempty"`
}
