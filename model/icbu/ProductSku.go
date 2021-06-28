package icbu

// ProductSku 
/* model for simplify = false
type ProductSku struct {

    // 商品属性
    
    Attributes  struct {
        ProductAttribute  []ProductAttribute `json:"product_attribute,omitempty"`
    } `json:"attributes,omitempty"`
    

    // 单个SKU详细定义
    
    SpecialSkus  struct {
        SkuDetail  []SkuDetail `json:"sku_detail,omitempty"`
    } `json:"special_skus,omitempty"`
    

    // 需要失效的SKU的详细定义
    
    ExcludeSkus  struct {
        SkuDetail  []SkuDetail `json:"sku_detail,omitempty"`
    } `json:"exclude_skus,omitempty"`
    

}
*/

// ProductSku 
type ProductSku struct {

    // 商品属性
    Attributes   []ProductAttribute `json:"attributes,omitempty"`

    // 单个SKU详细定义
    SpecialSkus   []SkuDetail `json:"special_skus,omitempty"`

    // 需要失效的SKU的详细定义
    ExcludeSkus   []SkuDetail `json:"exclude_skus,omitempty"`

}
