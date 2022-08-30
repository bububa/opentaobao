package idleisv

// YoupinCpvResultDto 结构体
type YoupinCpvResultDto struct {
	// 属性值list
	ValueList []YoupinPropertyValueResultDto `json:"value_list,omitempty" xml:"value_list>youpin_property_value_result_dto,omitempty"`
	// 属性id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 属性名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否包含子属性
	ShowSubProperty bool `json:"show_sub_property,omitempty" xml:"show_sub_property,omitempty"`
}
