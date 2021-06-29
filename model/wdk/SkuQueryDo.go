package wdk

// SkuQueryDo 
type SkuQueryDo struct {
    // 商品编码列表
    SkuCodes   []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
}
