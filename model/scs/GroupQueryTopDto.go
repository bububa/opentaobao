package scs

// GroupQueryTopDto 结构体
type GroupQueryTopDto struct {
	// 营销场景
	MarketScene string `json:"market_scene,omitempty" xml:"market_scene,omitempty"`
	// 模板ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 人群id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 页码
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}
