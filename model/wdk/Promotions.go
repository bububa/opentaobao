package wdk

import (
	"sync"
)

// Promotions 结构体
type Promotions struct {
	// 活动优惠金额
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolPromotions = sync.Pool{
	New: func() any {
		return new(Promotions)
	},
}

// GetPromotions() 从对象池中获取Promotions
func GetPromotions() *Promotions {
	return poolPromotions.Get().(*Promotions)
}

// ReleasePromotions 释放Promotions
func ReleasePromotions(v *Promotions) {
	v.DiscountFee = ""
	v.ActivityName = ""
	v.ActivityType = ""
	v.ActivityId = ""
	poolPromotions.Put(v)
}
