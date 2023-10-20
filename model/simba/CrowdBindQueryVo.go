package simba

import (
	"sync"
)

// CrowdBindQueryVo 结构体
type CrowdBindQueryVo struct {
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 单元id集合,单元已经存在场景必填
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 计划id,计划已经存在场景必填
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolCrowdBindQueryVo = sync.Pool{
	New: func() any {
		return new(CrowdBindQueryVo)
	},
}

// GetCrowdBindQueryVo() 从对象池中获取CrowdBindQueryVo
func GetCrowdBindQueryVo() *CrowdBindQueryVo {
	return poolCrowdBindQueryVo.Get().(*CrowdBindQueryVo)
}

// ReleaseCrowdBindQueryVo 释放CrowdBindQueryVo
func ReleaseCrowdBindQueryVo(v *CrowdBindQueryVo) {
	v.CampaignIdList = v.CampaignIdList[:0]
	v.AdgroupIdList = v.AdgroupIdList[:0]
	v.Offset = 0
	v.PageSize = 0
	v.CampaignId = 0
	poolCrowdBindQueryVo.Put(v)
}
