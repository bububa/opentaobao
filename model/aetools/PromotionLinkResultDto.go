package aetools

import (
	"sync"
)

// PromotionLinkResultDto 结构体
type PromotionLinkResultDto struct {
	// 推广链接列表
	PromotionLinks []PromotionLink `json:"promotion_links,omitempty" xml:"promotion_links>promotion_link,omitempty"`
	// 推广者TrackingId
	TrackingId string `json:"tracking_id,omitempty" xml:"tracking_id,omitempty"`
	// 返回总量
	TotalResultCount int64 `json:"total_result_count,omitempty" xml:"total_result_count,omitempty"`
}

var poolPromotionLinkResultDto = sync.Pool{
	New: func() any {
		return new(PromotionLinkResultDto)
	},
}

// GetPromotionLinkResultDto() 从对象池中获取PromotionLinkResultDto
func GetPromotionLinkResultDto() *PromotionLinkResultDto {
	return poolPromotionLinkResultDto.Get().(*PromotionLinkResultDto)
}

// ReleasePromotionLinkResultDto 释放PromotionLinkResultDto
func ReleasePromotionLinkResultDto(v *PromotionLinkResultDto) {
	v.PromotionLinks = v.PromotionLinks[:0]
	v.TrackingId = ""
	v.TotalResultCount = 0
	poolPromotionLinkResultDto.Put(v)
}
