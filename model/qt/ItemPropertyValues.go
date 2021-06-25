package qt

// ItemPropertyValues 
type ItemPropertyValues struct {

    // 服务属性id
    PropertyId   int64 `json:"property_id,omitempty"`

    // 属性名称
    PropertyName   string `json:"property_name,omitempty"`

    // 属性值列表.
    PropertyValues   []String `json:"property_values,omitempty"`

}
