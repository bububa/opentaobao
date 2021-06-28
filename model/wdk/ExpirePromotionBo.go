package wdk

// ExpirePromotionBo 
/* model for simplify = false
type ExpirePromotionBo struct {

    // 短保时间段信息
    
    PeriodInfos  struct {
        ExpirePeriodInfo  []ExpirePeriodInfo `json:"expire_period_info,omitempty"`
    } `json:"period_infos,omitempty"`
    

    // 门店id
    
    ShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"shop_ids,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 商品skucode
    
    SkuCode   string `json:"sku_code,omitempty"`
    

}
*/

// ExpirePromotionBo 
type ExpirePromotionBo struct {

    // 短保时间段信息
    PeriodInfos   []ExpirePeriodInfo `json:"period_infos,omitempty"`

    // 门店id
    ShopIds   []string `json:"shop_ids,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 商品skucode
    SkuCode   string `json:"sku_code,omitempty"`

}
