package xhotelonlineorder

import (
	"sync"
)

// TopOrderPromotionExtend 结构体
type TopOrderPromotionExtend struct {
	// topDailySellerPromotions
	TopDailySellerPromotions []TopDailySellerPromotion `json:"top_daily_seller_promotions,omitempty" xml:"top_daily_seller_promotions>top_daily_seller_promotion,omitempty"`
}

var poolTopOrderPromotionExtend = sync.Pool{
	New: func() any {
		return new(TopOrderPromotionExtend)
	},
}

// GetTopOrderPromotionExtend() 从对象池中获取TopOrderPromotionExtend
func GetTopOrderPromotionExtend() *TopOrderPromotionExtend {
	return poolTopOrderPromotionExtend.Get().(*TopOrderPromotionExtend)
}

// ReleaseTopOrderPromotionExtend 释放TopOrderPromotionExtend
func ReleaseTopOrderPromotionExtend(v *TopOrderPromotionExtend) {
	v.TopDailySellerPromotions = v.TopDailySellerPromotions[:0]
	poolTopOrderPromotionExtend.Put(v)
}
