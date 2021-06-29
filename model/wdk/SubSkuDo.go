package wdk

// SubSkuDo 
type SubSkuDo struct {
    // 子商品编码（需要先有子商品）
    SubSkuCode   string `json:"sub_sku_code,omitempty" xml:"sub_sku_code,omitempty"`
    // 子商品数量
    SubSkuNum   int64 `json:"sub_sku_num,omitempty" xml:"sub_sku_num,omitempty"`
}
