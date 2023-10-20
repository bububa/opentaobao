package feedflow

import (
	"sync"
)

// CrowdDto 结构体
type CrowdDto struct {
	// 人群描述
	CrowdDesc string `json:"crowd_desc,omitempty" xml:"crowd_desc,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 人群出价，单位：分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 定向
	TargetLabel *LabelDto `json:"target_label,omitempty" xml:"target_label,omitempty"`
	// 人群平均出价，单位：分
	AveragePrice int64 `json:"average_price,omitempty" xml:"average_price,omitempty"`
	// 人群建议出价，单位：分
	SuggestPrice int64 `json:"suggest_price,omitempty" xml:"suggest_price,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}

var poolCrowdDto = sync.Pool{
	New: func() any {
		return new(CrowdDto)
	},
}

// GetCrowdDto() 从对象池中获取CrowdDto
func GetCrowdDto() *CrowdDto {
	return poolCrowdDto.Get().(*CrowdDto)
}

// ReleaseCrowdDto 释放CrowdDto
func ReleaseCrowdDto(v *CrowdDto) {
	v.CrowdDesc = ""
	v.CrowdName = ""
	v.Status = ""
	v.Price = 0
	v.TargetLabel = nil
	v.AveragePrice = 0
	v.SuggestPrice = 0
	v.CrowdId = 0
	v.CampaignId = 0
	v.AdgroupId = 0
	poolCrowdDto.Put(v)
}
