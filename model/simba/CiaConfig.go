package simba

import (
	"sync"
)

// CiaConfig 结构体
type CiaConfig struct {
	// 目标点击量
	TargetClick int64 `json:"target_click,omitempty" xml:"target_click,omitempty"`
	// 出价偏好（3:促进收藏加购，4:促进点击，5:促进成交）
	BidTargetType int64 `json:"bid_target_type,omitempty" xml:"bid_target_type,omitempty"`
	// 最高溢价比例
	MaxPremium int64 `json:"max_premium,omitempty" xml:"max_premium,omitempty"`
	// 推广组id
	AdGroupId int64 `json:"ad_group_id,omitempty" xml:"ad_group_id,omitempty"`
	// 是否自动流转
	IsCirculation bool `json:"is_circulation,omitempty" xml:"is_circulation,omitempty"`
	// 是否开启智能出价
	IsSmartBidding bool `json:"is_smart_bidding,omitempty" xml:"is_smart_bidding,omitempty"`
}

var poolCiaConfig = sync.Pool{
	New: func() any {
		return new(CiaConfig)
	},
}

// GetCiaConfig() 从对象池中获取CiaConfig
func GetCiaConfig() *CiaConfig {
	return poolCiaConfig.Get().(*CiaConfig)
}

// ReleaseCiaConfig 释放CiaConfig
func ReleaseCiaConfig(v *CiaConfig) {
	v.TargetClick = 0
	v.BidTargetType = 0
	v.MaxPremium = 0
	v.AdGroupId = 0
	v.IsCirculation = false
	v.IsSmartBidding = false
	poolCiaConfig.Put(v)
}
