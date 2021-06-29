package promotion

// TaobaoUmpPromotionGlobalDiscountGetResult 
type TaobaoUmpPromotionGlobalDiscountGetResult struct {
    // 是否执行成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // defaultModel
    DefaultModel   *SellerGlobalDiscount `json:"default_model,omitempty" xml:"default_model,omitempty"`
}
