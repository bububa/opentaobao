package product

// MultiPropertyValueDto 结构体
type MultiPropertyValueDto struct {
	// 属性值对象数组
	PropertyValueList []PropertyValueDto `json:"property_value_list,omitempty" xml:"property_value_list>property_value_dto,omitempty"`
	// 父属性名
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
}
