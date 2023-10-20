package simba

import (
	"sync"
)

// CampaignBudget 结构体
type CampaignBudget struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 日限额，单位是元，不得小于30
	Budget int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// 是否平滑消耗，true-是；false-否；在设置了推广计划日限额后，此属性才生效
	IsSmooth bool `json:"is_smooth,omitempty" xml:"is_smooth,omitempty"`
}

var poolCampaignBudget = sync.Pool{
	New: func() any {
		return new(CampaignBudget)
	},
}

// GetCampaignBudget() 从对象池中获取CampaignBudget
func GetCampaignBudget() *CampaignBudget {
	return poolCampaignBudget.Get().(*CampaignBudget)
}

// ReleaseCampaignBudget 释放CampaignBudget
func ReleaseCampaignBudget(v *CampaignBudget) {
	v.Nick = ""
	v.CreateTime = ""
	v.ModifiedTime = ""
	v.CampaignId = 0
	v.Budget = 0
	v.IsSmooth = false
	poolCampaignBudget.Put(v)
}
