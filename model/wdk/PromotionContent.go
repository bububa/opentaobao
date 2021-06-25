package wdk

// PromotionContent 
type PromotionContent struct {

    // 客户商家编码
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty"`

    // sku列表
    PromotionSkuList   []PromotionSku `json:"promotion_sku_list,omitempty"`

    // 客户编码
    CustomerCode   string `json:"customer_code,omitempty"`

    // 客户门店
    OuCode   string `json:"ou_code,omitempty"`

    // 促销档期编码
    OuterPromotionCode   string `json:"outer_promotion_code,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 进售价类型
    PromotionType   string `json:"promotion_type,omitempty"`

}
