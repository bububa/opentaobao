package wdk

// SkuQueryDo 
/* model for simplify = false
type SkuQueryDo struct {

    // 商品编码列表
    
    SkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_codes,omitempty"`
    

    // 门店或DC编码
    
    OuCode   string `json:"ou_code,omitempty"`
    

}
*/

// SkuQueryDo 
type SkuQueryDo struct {

    // 商品编码列表
    SkuCodes   []string `json:"sku_codes,omitempty"`

    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty"`

}
