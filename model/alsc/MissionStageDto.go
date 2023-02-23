package alsc

// MissionStageDto 结构体
type MissionStageDto struct {
	// 奖品信息
	Rewards []Rewards `json:"rewards,omitempty" xml:"rewards>rewards,omitempty"`
	// 发奖状态
	RewardStatus string `json:"reward_status,omitempty" xml:"reward_status,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 阶段id
	StageCount int64 `json:"stage_count,omitempty" xml:"stage_count,omitempty"`
	// 任务配置阶段action
	SourceAction *MissionStageAction `json:"source_action,omitempty" xml:"source_action,omitempty"`
	// actionid
	SourceActionId int64 `json:"source_action_id,omitempty" xml:"source_action_id,omitempty"`
}
