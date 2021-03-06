package promotion

// TaobaoUmpPromotionGlobalDiscountGetResult 结构体
type TaobaoUmpPromotionGlobalDiscountGetResult struct {
	// defaultModel
	DefaultModel *SellerGlobalDiscount `json:"default_model,omitempty" xml:"default_model,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
