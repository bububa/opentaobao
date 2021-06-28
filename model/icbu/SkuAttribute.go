package icbu

// SkuAttribute 
type SkuAttribute struct {

    // 属性名称
    
    AttributeName   string `json:"attribute_name,omitempty" xml:"attribute_name,omitempty"`
    

    // 属性下的值
    
    Values   []SkuAttributeValue `json:"values,omitempty" xml:"values,omitempty"`
    

    // 属性ID
    
    AttributeId   int64 `json:"attribute_id,omitempty" xml:"attribute_id,omitempty"`
    

}
