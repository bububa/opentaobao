package simba

import (
	"sync"
)

// CrowdRefVo 结构体
type CrowdRefVo struct {
	// 横向管理
	ShowTagList []ShowTagVo `json:"show_tag_list,omitempty" xml:"show_tag_list>show_tag_vo,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 人群主键id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 定向状态,0:下线 1:上线
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 人群
	Crowd *CrowdVo `json:"crowd,omitempty" xml:"crowd,omitempty"`
	// 出价
	Price *PriceVo `json:"price,omitempty" xml:"price,omitempty"`
}

var poolCrowdRefVo = sync.Pool{
	New: func() any {
		return new(CrowdRefVo)
	},
}

// GetCrowdRefVo() 从对象池中获取CrowdRefVo
func GetCrowdRefVo() *CrowdRefVo {
	return poolCrowdRefVo.Get().(*CrowdRefVo)
}

// ReleaseCrowdRefVo 释放CrowdRefVo
func ReleaseCrowdRefVo(v *CrowdRefVo) {
	v.ShowTagList = v.ShowTagList[:0]
	v.CampaignId = 0
	v.AdgroupId = 0
	v.CrowdId = 0
	v.OnlineStatus = 0
	v.Crowd = nil
	v.Price = nil
	poolCrowdRefVo.Put(v)
}
