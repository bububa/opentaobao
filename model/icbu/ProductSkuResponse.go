package icbu

// ProductSkuResponse 
/* model for simplify = false
type ProductSkuResponse struct {

    // SKU使用的属性
    
    SkuAttributes  struct {
        SkuAttribute  []SkuAttribute `json:"sku_attribute,omitempty"`
    } `json:"sku_attributes,omitempty"`
    

    // SKU定义
    
    Skus  struct {
        SkuDefinition  []SkuDefinition `json:"sku_definition,omitempty"`
    } `json:"skus,omitempty"`
    

}
*/

// ProductSkuResponse 
type ProductSkuResponse struct {

    // SKU使用的属性
    SkuAttributes   []SkuAttribute `json:"sku_attributes,omitempty"`

    // SKU定义
    Skus   []SkuDefinition `json:"skus,omitempty"`

}
