package scs

// SolutionTopDto 结构体
type SolutionTopDto struct {
	// 计划信息
	CampaignView *CampaignSolutionTopDto `json:"campaign_view,omitempty" xml:"campaign_view,omitempty"`
}
