package alsc

// MissionDto 结构体
type MissionDto struct {
	// 任务阶段
	MissionStageDTOS []MissionStageDto `json:"mission_stage_d_t_o_s,omitempty" xml:"mission_stage_d_t_o_s>mission_stage_dto,omitempty"`
	// 任务xid
	MissionXId string `json:"mission_x_id,omitempty" xml:"mission_x_id,omitempty"`
	// 任务领取状态
	ReceiveStatus string `json:"receive_status,omitempty" xml:"receive_status,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 展示标题
	ShowTitle string `json:"show_title,omitempty" xml:"show_title,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 领取类型
	ReceiveType string `json:"receive_type,omitempty" xml:"receive_type,omitempty"`
	// 任务实例状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 扩展信息
	Ext *MissionExt `json:"ext,omitempty" xml:"ext,omitempty"`
	// 任务定义id
	MissionDefId int64 `json:"mission_def_id,omitempty" xml:"mission_def_id,omitempty"`
	// 下一个阶段
	NextStageCount int64 `json:"next_stage_count,omitempty" xml:"next_stage_count,omitempty"`
	// 行为配置
	ActionConfig *MissionAction `json:"action_config,omitempty" xml:"action_config,omitempty"`
	// 任务集id
	MissionCollectionId int64 `json:"mission_collection_id,omitempty" xml:"mission_collection_id,omitempty"`
	// 任务周期配置
	CycleConfig *MissionCycle `json:"cycle_config,omitempty" xml:"cycle_config,omitempty"`
}
