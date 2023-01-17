package alsc

// MissionStageAction 结构体
type MissionStageAction struct {
	// 阶段类型
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 阶段子类型
	ActionSubtype string `json:"action_subtype,omitempty" xml:"action_subtype,omitempty"`
	// 阶段子类型子类型
	ActionSubSubType string `json:"action_sub_sub_type,omitempty" xml:"action_sub_sub_type,omitempty"`
	// 奖励配置参数
	ActionParam string `json:"action_param,omitempty" xml:"action_param,omitempty"`
	//  阶段限制汇总
	StageSum string `json:"stage_sum,omitempty" xml:"stage_sum,omitempty"`
	// 阶段数
	StageCount int64 `json:"stage_count,omitempty" xml:"stage_count,omitempty"`
	// 阶段id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
