package scbp

import (
	"sync"
)

// TopP4pCampaignTargetingTagView 结构体
type TopP4pCampaignTargetingTagView struct {
	// 国家溢价列表
	CountryTagList []CountryTagView `json:"country_tag_list,omitempty" xml:"country_tag_list>country_tag_view,omitempty"`
	// 人群溢价列表
	CrowdTagList []CrowdTagView `json:"crowd_tag_list,omitempty" xml:"crowd_tag_list>crowd_tag_view,omitempty"`
}

var poolTopP4pCampaignTargetingTagView = sync.Pool{
	New: func() any {
		return new(TopP4pCampaignTargetingTagView)
	},
}

// GetTopP4pCampaignTargetingTagView() 从对象池中获取TopP4pCampaignTargetingTagView
func GetTopP4pCampaignTargetingTagView() *TopP4pCampaignTargetingTagView {
	return poolTopP4pCampaignTargetingTagView.Get().(*TopP4pCampaignTargetingTagView)
}

// ReleaseTopP4pCampaignTargetingTagView 释放TopP4pCampaignTargetingTagView
func ReleaseTopP4pCampaignTargetingTagView(v *TopP4pCampaignTargetingTagView) {
	v.CountryTagList = v.CountryTagList[:0]
	v.CrowdTagList = v.CrowdTagList[:0]
	poolTopP4pCampaignTargetingTagView.Put(v)
}
