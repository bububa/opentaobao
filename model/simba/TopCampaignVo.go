package simba

import (
	"sync"
)

// TopCampaignVo 结构体
type TopCampaignVo struct {
	// 计划实体
	Campaign *CampaignVo `json:"campaign,omitempty" xml:"campaign,omitempty"`
}

var poolTopCampaignVo = sync.Pool{
	New: func() any {
		return new(TopCampaignVo)
	},
}

// GetTopCampaignVo() 从对象池中获取TopCampaignVo
func GetTopCampaignVo() *TopCampaignVo {
	return poolTopCampaignVo.Get().(*TopCampaignVo)
}

// ReleaseTopCampaignVo 释放TopCampaignVo
func ReleaseTopCampaignVo(v *TopCampaignVo) {
	v.Campaign = nil
	poolTopCampaignVo.Put(v)
}
