package xhotelonlineorder

import (
	"sync"
)

// TopDailySellerPromotion 结构体
type TopDailySellerPromotion struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 卖家优惠金额
	SellerPromotion int64 `json:"seller_promotion,omitempty" xml:"seller_promotion,omitempty"`
}

var poolTopDailySellerPromotion = sync.Pool{
	New: func() any {
		return new(TopDailySellerPromotion)
	},
}

// GetTopDailySellerPromotion() 从对象池中获取TopDailySellerPromotion
func GetTopDailySellerPromotion() *TopDailySellerPromotion {
	return poolTopDailySellerPromotion.Get().(*TopDailySellerPromotion)
}

// ReleaseTopDailySellerPromotion 释放TopDailySellerPromotion
func ReleaseTopDailySellerPromotion(v *TopDailySellerPromotion) {
	v.Date = ""
	v.SellerPromotion = 0
	poolTopDailySellerPromotion.Put(v)
}
