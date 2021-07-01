package xhotelonlineorder

// TopDailySellerPromotion 结构体
type TopDailySellerPromotion struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 卖家优惠金额
	SellerPromotion int64 `json:"seller_promotion,omitempty" xml:"seller_promotion,omitempty"`
}
