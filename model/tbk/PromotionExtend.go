package tbk

// PromotionExtend 结构体
type PromotionExtend struct {
	// 权益推荐商品
	RecommendItemList []RecommendItemList `json:"recommend_item_list,omitempty" xml:"recommend_item_list>recommend_item_list,omitempty"`
	// 权益链接
	PromotionUrl string `json:"promotion_url,omitempty" xml:"promotion_url,omitempty"`
	// 有价券信息
	YoujiaCouponInfo *Youjiacouponinfo `json:"youjia_coupon_info,omitempty" xml:"youjia_coupon_info,omitempty"`
}
