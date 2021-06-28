package promotion

// TaobaoUmpPromotionGlobalDiscountGetResult 
/* model for simplify = false
type TaobaoUmpPromotionGlobalDiscountGetResult struct {

    // 是否执行成功
    
    Success   bool `json:"success,omitempty"`
    

    // defaultModel
    
    DefaultModel  *struct {
        SellerGlobalDiscount  *SellerGlobalDiscount `json:"seller_global_discount,omitempty"`
    } `json:"default_model,omitempty"`
    

}
*/

// TaobaoUmpPromotionGlobalDiscountGetResult 
type TaobaoUmpPromotionGlobalDiscountGetResult struct {

    // 是否执行成功
    Success   bool `json:"success,omitempty"`

    // defaultModel
    DefaultModel   *SellerGlobalDiscount `json:"default_model,omitempty"`

}
