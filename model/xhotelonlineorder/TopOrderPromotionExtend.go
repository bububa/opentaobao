package xhotelonlineorder

// TopOrderPromotionExtend 结构体
type TopOrderPromotionExtend struct {
	// topDailySellerPromotions
	TopDailySellerPromotions []TopDailySellerPromotion `json:"top_daily_seller_promotions,omitempty" xml:"top_daily_seller_promotions>top_daily_seller_promotion,omitempty"`
}
