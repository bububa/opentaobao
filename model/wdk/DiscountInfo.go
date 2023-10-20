package wdk

import (
	"sync"
)

// DiscountInfo 结构体
type DiscountInfo struct {
	// 营销活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 营销活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动优惠金额商家分摊
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 活动优惠金额平台分摊
	PlatformDiscountFee int64 `json:"platform_discount_fee,omitempty" xml:"platform_discount_fee,omitempty"`
	// 优惠金额
	DicountFee int64 `json:"dicount_fee,omitempty" xml:"dicount_fee,omitempty"`
}

var poolDiscountInfo = sync.Pool{
	New: func() any {
		return new(DiscountInfo)
	},
}

// GetDiscountInfo() 从对象池中获取DiscountInfo
func GetDiscountInfo() *DiscountInfo {
	return poolDiscountInfo.Get().(*DiscountInfo)
}

// ReleaseDiscountInfo 释放DiscountInfo
func ReleaseDiscountInfo(v *DiscountInfo) {
	v.ActivityId = ""
	v.ActivityType = ""
	v.ActivityName = ""
	v.DiscountFee = 0
	v.MerchantDiscountFee = 0
	v.PlatformDiscountFee = 0
	v.DicountFee = 0
	poolDiscountInfo.Put(v)
}
