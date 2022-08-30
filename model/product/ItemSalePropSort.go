package product

// ItemSalePropSort 结构体
type ItemSalePropSort struct {
	// 属性值列表
	SalePropValueSorts []SalePropValueSort `json:"sale_prop_value_sorts,omitempty" xml:"sale_prop_value_sorts>sale_prop_value_sort,omitempty"`
	// 属性项名称
	PropertyValue string `json:"property_value,omitempty" xml:"property_value,omitempty"`
	// 属性项ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
