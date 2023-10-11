package simba

// CampaignLaunchTimeVo 结构体
type CampaignLaunchTimeVo struct {
	// 计划开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 是否长期投放,true:是,false:否
	LaunchForever bool `json:"launch_forever,omitempty" xml:"launch_forever,omitempty"`
}
