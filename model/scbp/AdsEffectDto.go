package scbp

import (
	"sync"
)

// AdsEffectDto 结构体
type AdsEffectDto struct {
	// 曝光
	Impr int64 `json:"impr,omitempty" xml:"impr,omitempty"`
	// 点击
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 消耗
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 推广时长
	OnlineMin int64 `json:"online_min,omitempty" xml:"online_min,omitempty"`
}

var poolAdsEffectDto = sync.Pool{
	New: func() any {
		return new(AdsEffectDto)
	},
}

// GetAdsEffectDto() 从对象池中获取AdsEffectDto
func GetAdsEffectDto() *AdsEffectDto {
	return poolAdsEffectDto.Get().(*AdsEffectDto)
}

// ReleaseAdsEffectDto 释放AdsEffectDto
func ReleaseAdsEffectDto(v *AdsEffectDto) {
	v.Impr = 0
	v.Click = 0
	v.Cost = 0
	v.OnlineMin = 0
	poolAdsEffectDto.Put(v)
}
