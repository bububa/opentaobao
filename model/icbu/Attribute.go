package icbu

// Attribute 
type Attribute struct {

    // 属性id
    
    AttrId   int64 `json:"attr_id,omitempty" xml:"attr_id,omitempty"`
    

    // 英文名字
    
    EnName   string `json:"en_name,omitempty" xml:"en_name,omitempty"`
    

    // 展示类型；input；group
    
    ShowType   string `json:"show_type,omitempty" xml:"show_type,omitempty"`
    

    // 是否必填属性
    
    Required   bool `json:"required,omitempty" xml:"required,omitempty"`
    

    // 该属性的单位
    
    Units   []string `json:"units,omitempty" xml:"units>string,omitempty"`
    

    // 该属性能否当成SKU属性
    
    SkuAttribute   bool `json:"sku_attribute,omitempty" xml:"sku_attribute,omitempty"`
    

    // 用成SKU属性时，是否支持自定义图片展示
    
    CustomizeImage   bool `json:"customize_image,omitempty" xml:"customize_image,omitempty"`
    

    // 属性可选的属性值
    
    AttributeValues   []AttributeValue `json:"attribute_values,omitempty" xml:"attribute_values,omitempty"`
    

    // 输入类型
    
    InputType   string `json:"input_type,omitempty" xml:"input_type,omitempty"`
    

    // 用成SKU属性时，是否支持自定义属性值名称
    
    CustomizeValue   bool `json:"customize_value,omitempty" xml:"customize_value,omitempty"`
    

    // valueType
    
    ValueType   string `json:"value_type,omitempty" xml:"value_type,omitempty"`
    

    // 表示是否车型库属性，如果是，则需要从分层属性接口里获取下一级属性
    
    CarModel   bool `json:"car_model,omitempty" xml:"car_model,omitempty"`
    

}
