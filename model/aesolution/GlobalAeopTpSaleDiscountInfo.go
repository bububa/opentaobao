package aesolution

// GlobalAeopTpSaleDiscountInfo 
type GlobalAeopTpSaleDiscountInfo struct {
    // discount detail
    DiscountDetail   *GlobalMoneyStr `json:"discount_detail,omitempty" xml:"discount_detail,omitempty"`
    // Promotion owner, can be Seller, Platform or Mix.
    PromotionOwner   string `json:"promotion_owner,omitempty" xml:"promotion_owner,omitempty"`
    // promotion type
    PromotionType   string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
    // promotion ID
    PromotionId   string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
}
