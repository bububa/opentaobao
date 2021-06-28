package tbk

// PromotionExtend 
/* model for simplify = false
type PromotionExtend struct {

    // 权益推荐商品
    
    RecommendItemList  struct {
        RecommendItemList  []RecommendItemList `json:"recommend_item_list,omitempty"`
    } `json:"recommend_item_list,omitempty"`
    

    // 有价券信息
    
    YoujiaCouponInfo  *struct {
        Youjiacouponinfo  *Youjiacouponinfo `json:"youjiacouponinfo,omitempty"`
    } `json:"youjia_coupon_info,omitempty"`
    

    // 权益链接
    
    PromotionUrl   string `json:"promotion_url,omitempty"`
    

}
*/

// PromotionExtend 
type PromotionExtend struct {

    // 权益推荐商品
    RecommendItemList   []RecommendItemList `json:"recommend_item_list,omitempty"`

    // 有价券信息
    YoujiaCouponInfo   *Youjiacouponinfo `json:"youjia_coupon_info,omitempty"`

    // 权益链接
    PromotionUrl   string `json:"promotion_url,omitempty"`

}
