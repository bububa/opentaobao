package qt

// ItemPropertyValues 结构体
type ItemPropertyValues struct {
	// 服务属性id
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 属性名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 属性值列表.
	PropertyValues []string `json:"property_values,omitempty" xml:"property_values>string,omitempty"`
}
