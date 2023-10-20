package util

import (
	"sync"
)

// CrowdRefTopDto 结构体
type CrowdRefTopDto struct {
	// 人群信息
	Crowd *CrowdTopDto `json:"crowd,omitempty" xml:"crowd,omitempty"`
	// 溢价相关信息
	PriceTopDTO *PriceTopDto `json:"price_top_d_t_o,omitempty" xml:"price_top_d_t_o,omitempty"`
	// 需要添加到的计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 推广传1，暂停传0
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 需要添加到的单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}

var poolCrowdRefTopDto = sync.Pool{
	New: func() any {
		return new(CrowdRefTopDto)
	},
}

// GetCrowdRefTopDto() 从对象池中获取CrowdRefTopDto
func GetCrowdRefTopDto() *CrowdRefTopDto {
	return poolCrowdRefTopDto.Get().(*CrowdRefTopDto)
}

// ReleaseCrowdRefTopDto 释放CrowdRefTopDto
func ReleaseCrowdRefTopDto(v *CrowdRefTopDto) {
	v.Crowd = nil
	v.PriceTopDTO = nil
	v.CampaignId = 0
	v.OnlineStatus = 0
	v.AdgroupId = 0
	poolCrowdRefTopDto.Put(v)
}
