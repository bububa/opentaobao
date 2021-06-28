package icbu

// ProductAttribute 
/* model for simplify = false
type ProductAttribute struct {

    // 属性ID
    
    AttributeId   int64 `json:"attribute_id,omitempty"`
    

    // 属性名称
    
    AttributeName   string `json:"attribute_name,omitempty"`
    

    // 属性值ID
    
    ValueId   int64 `json:"value_id,omitempty"`
    

    // 属性值名称
    
    ValueName   string `json:"value_name,omitempty"`
    

    // 作为sku属性值时，自定义属性值名称
    
    SkuCustomValueName   string `json:"sku_custom_value_name,omitempty"`
    

    // 作为sku属性值时，用图形来展示；必须是alibaba图片中心的图片URL，请使用API alibaba.icbu.photobank.upload 上传图片
    
    SkuCustomImageUrl   string `json:"sku_custom_image_url,omitempty"`
    

}
*/

// ProductAttribute 
type ProductAttribute struct {

    // 属性ID
    AttributeId   int64 `json:"attribute_id,omitempty"`

    // 属性名称
    AttributeName   string `json:"attribute_name,omitempty"`

    // 属性值ID
    ValueId   int64 `json:"value_id,omitempty"`

    // 属性值名称
    ValueName   string `json:"value_name,omitempty"`

    // 作为sku属性值时，自定义属性值名称
    SkuCustomValueName   string `json:"sku_custom_value_name,omitempty"`

    // 作为sku属性值时，用图形来展示；必须是alibaba图片中心的图片URL，请使用API alibaba.icbu.photobank.upload 上传图片
    SkuCustomImageUrl   string `json:"sku_custom_image_url,omitempty"`

}
