package flight

// ConfirmRequestDto 结构体
type ConfirmRequestDto struct {
	// 商家店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 意向单id
	IntentionId int64 `json:"intention_id,omitempty" xml:"intention_id,omitempty"`
}
