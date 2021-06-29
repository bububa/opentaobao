package xhotelonlineorder

// TopPromotion 
type TopPromotion struct {

    // topCreateOrderPromotion
    
    TopCreateOrderPromotion   *TopOrderPromotionDetail `json:"top_create_order_promotion,omitempty" xml:"top_create_order_promotion,omitempty"`
    

    // topCurrentOrderPromotion
    
    TopCurrentOrderPromotion   *TopOrderPromotionDetail `json:"top_current_order_promotion,omitempty" xml:"top_current_order_promotion,omitempty"`
    

    // topOrderPromotionExtend
    
    TopOrderPromotionExtend   *TopOrderPromotionExtend `json:"top_order_promotion_extend,omitempty" xml:"top_order_promotion_extend,omitempty"`
    

}
