package icbu

// ProductSku 
type ProductSku struct {

    // 商品属性
    
    Attributes   []ProductAttribute `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 单个SKU详细定义
    
    SpecialSkus   []SkuDetail `json:"special_skus,omitempty" xml:"special_skus,omitempty"`
    

    // 需要失效的SKU的详细定义
    
    ExcludeSkus   []SkuDetail `json:"exclude_skus,omitempty" xml:"exclude_skus,omitempty"`
    

}
