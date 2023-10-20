package singletreasure

import (
	"sync"
)

// ActivityInfo 结构体
type ActivityInfo struct {
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 不包邮地区
	ExcludeAreas string `json:"exclude_areas,omitempty" xml:"exclude_areas,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 套餐描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 套餐名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动对应生效的人群 id
	CrowdId string `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 套餐 id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 折扣类型，1 为促销价，2 为打折，4 为减钱
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 卖家 id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 优惠类型，3 为 SKU 级优惠类型，1 为商品级优惠类型
	PromotionLevel int64 `json:"promotion_level,omitempty" xml:"promotion_level,omitempty"`
	// 活动状态
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
	// 是否包邮
	FreePost bool `json:"free_post,omitempty" xml:"free_post,omitempty"`
}

var poolActivityInfo = sync.Pool{
	New: func() any {
		return new(ActivityInfo)
	},
}

// GetActivityInfo() 从对象池中获取ActivityInfo
func GetActivityInfo() *ActivityInfo {
	return poolActivityInfo.Get().(*ActivityInfo)
}

// ReleaseActivityInfo 释放ActivityInfo
func ReleaseActivityInfo(v *ActivityInfo) {
	v.CreatedTime = ""
	v.EndTime = ""
	v.ExcludeAreas = ""
	v.StartTime = ""
	v.Description = ""
	v.Name = ""
	v.CrowdId = ""
	v.ActivityId = 0
	v.DiscountType = 0
	v.SellerId = 0
	v.PromotionLevel = 0
	v.ActivityStatus = 0
	v.FreePost = false
	poolActivityInfo.Put(v)
}
