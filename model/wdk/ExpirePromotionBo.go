package wdk

// ExpirePromotionBo 
type ExpirePromotionBo struct {

    // 短保时间段信息
    PeriodInfos   []ExpirePeriodInfo `json:"period_infos,omitempty"`

    // 门店id
    ShopIds   []String `json:"shop_ids,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 商品skucode
    SkuCode   string `json:"sku_code,omitempty"`

}
