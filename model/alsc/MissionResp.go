package alsc

// MissionResp 结构体
type MissionResp struct {
	// 任务列表
	MList []MissionDto `json:"m_list,omitempty" xml:"m_list>mission_dto,omitempty"`
	// 提醒信息
	RemindInfo *RemindInfo `json:"remind_info,omitempty" xml:"remind_info,omitempty"`
	// 任务集配置信息
	CollectionDto *CollectionDto `json:"collection_dto,omitempty" xml:"collection_dto,omitempty"`
}
