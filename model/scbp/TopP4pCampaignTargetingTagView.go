package scbp

// TopP4pCampaignTargetingTagView 结构体
type TopP4pCampaignTargetingTagView struct {
	// 国家溢价列表
	CountryTagList []CountryTagView `json:"country_tag_list,omitempty" xml:"country_tag_list>country_tag_view,omitempty"`
	// 人群溢价列表
	CrowdTagList []CrowdTagView `json:"crowd_tag_list,omitempty" xml:"crowd_tag_list>crowd_tag_view,omitempty"`
}
