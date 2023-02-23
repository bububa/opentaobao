package alsc

// MissionAction 结构体
type MissionAction struct {
	// 行为类型
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 行为配置
	ActionValue string `json:"action_value,omitempty" xml:"action_value,omitempty"`
}
