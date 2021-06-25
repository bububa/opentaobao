package tbk

// PromotionExtend 
type PromotionExtend struct {

    // 权益推荐商品
    RecommendItemList   []RecommendItemList `json:"recommend_item_list,omitempty"`

    // 有价券信息
    YoujiaCouponInfo   *Youjiacouponinfo `json:"youjia_coupon_info,omitempty"`

    // 权益链接
    PromotionUrl   string `json:"promotion_url,omitempty"`

}
