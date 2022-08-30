package usergrowth2

// OfflinePromoteInfo 结构体
type OfflinePromoteInfo struct {
	// 新登数据明细
	TransformInfos []OfflineUserTransformInfo `json:"transform_infos,omitempty" xml:"transform_infos>offline_user_transform_info,omitempty"`
	// 推广码
	PromoterCode string `json:"promoter_code,omitempty" xml:"promoter_code,omitempty"`
	// 任务名
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 活动名
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 渠道id
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
