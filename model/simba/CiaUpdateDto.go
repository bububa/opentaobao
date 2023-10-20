package simba

import (
	"sync"
)

// CiaUpdateDto 结构体
type CiaUpdateDto struct {
	// 计划Id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 出价目标（3:促进收藏加购,4:促进点击,5:促进成交）
	BidTargetType int64 `json:"bid_target_type,omitempty" xml:"bid_target_type,omitempty"`
	// 最高溢价比例（30-300）
	MaxPremium int64 `json:"max_premium,omitempty" xml:"max_premium,omitempty"`
	// 是否自动流转（0:否，1:是）
	IsCirculation int64 `json:"is_circulation,omitempty" xml:"is_circulation,omitempty"`
	// 推广组Id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 是否开启智能出价
	IsSmartBidding bool `json:"is_smart_bidding,omitempty" xml:"is_smart_bidding,omitempty"`
}

var poolCiaUpdateDto = sync.Pool{
	New: func() any {
		return new(CiaUpdateDto)
	},
}

// GetCiaUpdateDto() 从对象池中获取CiaUpdateDto
func GetCiaUpdateDto() *CiaUpdateDto {
	return poolCiaUpdateDto.Get().(*CiaUpdateDto)
}

// ReleaseCiaUpdateDto 释放CiaUpdateDto
func ReleaseCiaUpdateDto(v *CiaUpdateDto) {
	v.CampaignId = 0
	v.BidTargetType = 0
	v.MaxPremium = 0
	v.IsCirculation = 0
	v.AdgroupId = 0
	v.IsSmartBidding = false
	poolCiaUpdateDto.Put(v)
}
