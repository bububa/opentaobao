package singletreasure

import (
	"sync"
)

// ActivityInfoCreateDto 结构体
type ActivityInfoCreateDto struct {
	// 开始时间：需要早于当天零点
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 人群 id
	CrowdId string `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 不包邮地区
	ExcludeAreas string `json:"exclude_areas,omitempty" xml:"exclude_areas,omitempty"`
	// 优惠类型    促销价：1；打折：2；减钱：4
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 优惠等级 商品:1; 店铺:2; SKU:3
	PromotionLevel int64 `json:"promotion_level,omitempty" xml:"promotion_level,omitempty"`
	// 活动类型, 普通活动: 1, 官方活动: 2
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动的Id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否包邮
	FreePost bool `json:"free_post,omitempty" xml:"free_post,omitempty"`
}

var poolActivityInfoCreateDto = sync.Pool{
	New: func() any {
		return new(ActivityInfoCreateDto)
	},
}

// GetActivityInfoCreateDto() 从对象池中获取ActivityInfoCreateDto
func GetActivityInfoCreateDto() *ActivityInfoCreateDto {
	return poolActivityInfoCreateDto.Get().(*ActivityInfoCreateDto)
}

// ReleaseActivityInfoCreateDto 释放ActivityInfoCreateDto
func ReleaseActivityInfoCreateDto(v *ActivityInfoCreateDto) {
	v.StartTime = ""
	v.Description = ""
	v.Name = ""
	v.CrowdId = ""
	v.EndTime = ""
	v.ExcludeAreas = ""
	v.DiscountType = 0
	v.PromotionLevel = 0
	v.ActivityType = 0
	v.ActivityId = 0
	v.FreePost = false
	poolActivityInfoCreateDto.Put(v)
}
