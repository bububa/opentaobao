package alitripbp

// ChannelExamineResultDto 结构体
type ChannelExamineResultDto struct {
	// 活动url
	TargetUrl string `json:"target_url,omitempty" xml:"target_url,omitempty"`
	// 是否为活动的目标用户
	IsTargetCrow bool `json:"is_target_crow,omitempty" xml:"is_target_crow,omitempty"`
}
