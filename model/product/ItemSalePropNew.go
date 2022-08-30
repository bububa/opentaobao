package product

// ItemSalePropNew 结构体
type ItemSalePropNew struct {
	// 属性状态集合
	SalePropValueStatusList []SalePropValueStatus `json:"sale_prop_value_status_list,omitempty" xml:"sale_prop_value_status_list>sale_prop_value_status,omitempty"`
	// 属性文本
	PropertyValue string `json:"property_value,omitempty" xml:"property_value,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
