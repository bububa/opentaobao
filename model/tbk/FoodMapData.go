package tbk

// FoodMapData 结构体
type FoodMapData struct {
	// 本地化-商品图片
	FoodPic string `json:"food_pic,omitempty" xml:"food_pic,omitempty"`
	// 本地化-商品标题
	FoodTitle string `json:"food_title,omitempty" xml:"food_title,omitempty"`
	// 本地化-商品促销价
	FoodPromotionPrice string `json:"food_promotion_price,omitempty" xml:"food_promotion_price,omitempty"`
	// 本地化-商品原价
	FoodReservePrice string `json:"food_reserve_price,omitempty" xml:"food_reserve_price,omitempty"`
}
