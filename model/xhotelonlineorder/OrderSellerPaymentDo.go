package xhotelonlineorder

import (
	"sync"
)

// OrderSellerPaymentDo 结构体
type OrderSellerPaymentDo struct {
	// 下单时刻营销后金额
	PromotionCreatePrice int64 `json:"promotion_create_price,omitempty" xml:"promotion_create_price,omitempty"`
	// 下单时刻营销后金额
	AfterPromotionCreatePrice int64 `json:"after_promotion_create_price,omitempty" xml:"after_promotion_create_price,omitempty"`
	// 总房费
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 每日价格
	DailyInfos *HbsDailyPriceDo `json:"daily_infos,omitempty" xml:"daily_infos,omitempty"`
	// 优惠后外币金额
	AfterPromotionCurrentPrice int64 `json:"after_promotion_current_price,omitempty" xml:"after_promotion_current_price,omitempty"`
	// 优惠外币金额
	PromotionCurrentPrice int64 `json:"promotion_current_price,omitempty" xml:"promotion_current_price,omitempty"`
}

var poolOrderSellerPaymentDo = sync.Pool{
	New: func() any {
		return new(OrderSellerPaymentDo)
	},
}

// GetOrderSellerPaymentDo() 从对象池中获取OrderSellerPaymentDo
func GetOrderSellerPaymentDo() *OrderSellerPaymentDo {
	return poolOrderSellerPaymentDo.Get().(*OrderSellerPaymentDo)
}

// ReleaseOrderSellerPaymentDo 释放OrderSellerPaymentDo
func ReleaseOrderSellerPaymentDo(v *OrderSellerPaymentDo) {
	v.PromotionCreatePrice = 0
	v.AfterPromotionCreatePrice = 0
	v.TotalRoomPrice = 0
	v.DailyInfos = nil
	v.AfterPromotionCurrentPrice = 0
	v.PromotionCurrentPrice = 0
	poolOrderSellerPaymentDo.Put(v)
}
