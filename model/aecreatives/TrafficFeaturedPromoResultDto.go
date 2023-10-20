package aecreatives

import (
	"sync"
)

// TrafficFeaturedPromoResultDto 结构体
type TrafficFeaturedPromoResultDto struct {
	// 返回主题活动列表
	Promos []Promo `json:"promos,omitempty" xml:"promos>promo,omitempty"`
	// 当前返回数量
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
}

var poolTrafficFeaturedPromoResultDto = sync.Pool{
	New: func() any {
		return new(TrafficFeaturedPromoResultDto)
	},
}

// GetTrafficFeaturedPromoResultDto() 从对象池中获取TrafficFeaturedPromoResultDto
func GetTrafficFeaturedPromoResultDto() *TrafficFeaturedPromoResultDto {
	return poolTrafficFeaturedPromoResultDto.Get().(*TrafficFeaturedPromoResultDto)
}

// ReleaseTrafficFeaturedPromoResultDto 释放TrafficFeaturedPromoResultDto
func ReleaseTrafficFeaturedPromoResultDto(v *TrafficFeaturedPromoResultDto) {
	v.Promos = v.Promos[:0]
	v.CurrentRecordCount = 0
	poolTrafficFeaturedPromoResultDto.Put(v)
}
