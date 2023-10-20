package xhotelonlineorder

import (
	"sync"
)

// TopOrderPromotionDetail 结构体
type TopOrderPromotionDetail struct {
	// topHotelPromotions
	TopHotelPromotions []TopHotelPromotion `json:"top_hotel_promotions,omitempty" xml:"top_hotel_promotions>top_hotel_promotion,omitempty"`
	// 返现金额
	CashBackAmount int64 `json:"cash_back_amount,omitempty" xml:"cash_back_amount,omitempty"`
	// 卖家立减
	SellerDecrease int64 `json:"seller_decrease,omitempty" xml:"seller_decrease,omitempty"`
	// 立减金额
	Decrease int64 `json:"decrease,omitempty" xml:"decrease,omitempty"`
	// 平台立减金额
	PlatformDecrease int64 `json:"platform_decrease,omitempty" xml:"platform_decrease,omitempty"`
}

var poolTopOrderPromotionDetail = sync.Pool{
	New: func() any {
		return new(TopOrderPromotionDetail)
	},
}

// GetTopOrderPromotionDetail() 从对象池中获取TopOrderPromotionDetail
func GetTopOrderPromotionDetail() *TopOrderPromotionDetail {
	return poolTopOrderPromotionDetail.Get().(*TopOrderPromotionDetail)
}

// ReleaseTopOrderPromotionDetail 释放TopOrderPromotionDetail
func ReleaseTopOrderPromotionDetail(v *TopOrderPromotionDetail) {
	v.TopHotelPromotions = v.TopHotelPromotions[:0]
	v.CashBackAmount = 0
	v.SellerDecrease = 0
	v.Decrease = 0
	v.PlatformDecrease = 0
	poolTopOrderPromotionDetail.Put(v)
}
