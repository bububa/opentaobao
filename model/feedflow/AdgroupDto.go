package feedflow

import (
	"sync"
)

// AdgroupDto 结构体
type AdgroupDto struct {
	// 资源类位表
	AdzoneList []AdzoneBindDto `json:"adzone_list,omitempty" xml:"adzone_list>adzone_bind_dto,omitempty"`
	// 定向人群
	CrowdList []CrowdDto `json:"crowd_list,omitempty" xml:"crowd_list>crowd_dto,omitempty"`
	// 单元名称
	AdgroupName string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 智能调价
	IntelligentBid *IntelligentBidDto `json:"intelligent_bid,omitempty" xml:"intelligent_bid,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}

var poolAdgroupDto = sync.Pool{
	New: func() any {
		return new(AdgroupDto)
	},
}

// GetAdgroupDto() 从对象池中获取AdgroupDto
func GetAdgroupDto() *AdgroupDto {
	return poolAdgroupDto.Get().(*AdgroupDto)
}

// ReleaseAdgroupDto 释放AdgroupDto
func ReleaseAdgroupDto(v *AdgroupDto) {
	v.AdzoneList = v.AdzoneList[:0]
	v.CrowdList = v.CrowdList[:0]
	v.AdgroupName = ""
	v.Status = ""
	v.CampaignId = 0
	v.IntelligentBid = nil
	v.ItemId = 0
	v.AdgroupId = 0
	poolAdgroupDto.Put(v)
}
