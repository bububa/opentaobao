package xhotelonlineorder

import (
	"sync"
)

// OrderPromotionDetail 结构体
type OrderPromotionDetail struct {
	// 优惠日历数据，如{时间:{\&#34;amount\&#34;:每日优惠金额,},时间:{\&#34;amount\&#34;:每日优惠金额}}
	DailyHotelPromotion string `json:"daily_hotel_promotion,omitempty" xml:"daily_hotel_promotion,omitempty"`
	// 总卖家优惠
	SellerDecrease int64 `json:"seller_decrease,omitempty" xml:"seller_decrease,omitempty"`
}

var poolOrderPromotionDetail = sync.Pool{
	New: func() any {
		return new(OrderPromotionDetail)
	},
}

// GetOrderPromotionDetail() 从对象池中获取OrderPromotionDetail
func GetOrderPromotionDetail() *OrderPromotionDetail {
	return poolOrderPromotionDetail.Get().(*OrderPromotionDetail)
}

// ReleaseOrderPromotionDetail 释放OrderPromotionDetail
func ReleaseOrderPromotionDetail(v *OrderPromotionDetail) {
	v.DailyHotelPromotion = ""
	v.SellerDecrease = 0
	poolOrderPromotionDetail.Put(v)
}
