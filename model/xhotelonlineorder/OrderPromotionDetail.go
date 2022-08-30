package xhotelonlineorder

// OrderPromotionDetail 结构体
type OrderPromotionDetail struct {
	// 优惠日历数据，如{时间:{\"amount\":每日优惠金额,},时间:{\"amount\":每日优惠金额}}
	DailyHotelPromotion string `json:"daily_hotel_promotion,omitempty" xml:"daily_hotel_promotion,omitempty"`
	// 总卖家优惠
	SellerDecrease int64 `json:"seller_decrease,omitempty" xml:"seller_decrease,omitempty"`
}
