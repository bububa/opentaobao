package alsc

// RelationConfigDto 结构体
type RelationConfigDto struct {
	// 子任务列表
	ChildIds []int64 `json:"child_ids,omitempty" xml:"child_ids>int64,omitempty"`
	// 任务详细关系
	MissionRelation string `json:"mission_relation,omitempty" xml:"mission_relation,omitempty"`
	// 父子关系
	MissionRelationType string `json:"mission_relation_type,omitempty" xml:"mission_relation_type,omitempty"`
	// 关系名称
	RelationName string `json:"relation_name,omitempty" xml:"relation_name,omitempty"`
	// 父任务id
	FatherId int64 `json:"father_id,omitempty" xml:"father_id,omitempty"`
	// 子任务领取限制
	SubMissionReceiveCount int64 `json:"sub_mission_receive_count,omitempty" xml:"sub_mission_receive_count,omitempty"`
	// 子任务展示个数
	SubMissionShowCount int64 `json:"sub_mission_show_count,omitempty" xml:"sub_mission_show_count,omitempty"`
}
