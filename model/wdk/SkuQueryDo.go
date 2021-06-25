package wdk

// SkuQueryDo 
type SkuQueryDo struct {

    // 商品编码列表
    SkuCodes   []String `json:"sku_codes,omitempty"`

    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty"`

}