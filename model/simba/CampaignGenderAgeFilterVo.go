package simba

import (
	"sync"
)

// CampaignGenderAgeFilterVo 结构体
type CampaignGenderAgeFilterVo struct {
	// 年龄段,0:,1,2,3,4,5,6
	AgeList []string `json:"age_list,omitempty" xml:"age_list>string,omitempty"`
	// 性别,man:男性人群,woman:女性人群
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
}

var poolCampaignGenderAgeFilterVo = sync.Pool{
	New: func() any {
		return new(CampaignGenderAgeFilterVo)
	},
}

// GetCampaignGenderAgeFilterVo() 从对象池中获取CampaignGenderAgeFilterVo
func GetCampaignGenderAgeFilterVo() *CampaignGenderAgeFilterVo {
	return poolCampaignGenderAgeFilterVo.Get().(*CampaignGenderAgeFilterVo)
}

// ReleaseCampaignGenderAgeFilterVo 释放CampaignGenderAgeFilterVo
func ReleaseCampaignGenderAgeFilterVo(v *CampaignGenderAgeFilterVo) {
	v.AgeList = v.AgeList[:0]
	v.Gender = ""
	poolCampaignGenderAgeFilterVo.Put(v)
}
