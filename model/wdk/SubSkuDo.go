package wdk

// SubSkuDo 
/* model for simplify = false
type SubSkuDo struct {

    // 子商品编码（需要先有子商品）
    
    SubSkuCode   string `json:"sub_sku_code,omitempty"`
    

    // 子商品数量
    
    SubSkuNum   int64 `json:"sub_sku_num,omitempty"`
    

}
*/

// SubSkuDo 
type SubSkuDo struct {

    // 子商品编码（需要先有子商品）
    SubSkuCode   string `json:"sub_sku_code,omitempty"`

    // 子商品数量
    SubSkuNum   int64 `json:"sub_sku_num,omitempty"`

}
