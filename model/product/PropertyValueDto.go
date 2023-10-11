package product

// PropertyValueDto 结构体
type PropertyValueDto struct {
	// 子属性列表
	ChildPropertyKeyValueList []ChildPropertyKeyValueDto `json:"child_property_key_value_list,omitempty" xml:"child_property_key_value_list>child_property_key_value_dto,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
