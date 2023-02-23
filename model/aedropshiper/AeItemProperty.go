package aedropshiper

// AeItemProperty 结构体
type AeItemProperty struct {
	// Attribute name
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// Attribute value
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	// Interval attribute start value
	AttrValueStart string `json:"attr_value_start,omitempty" xml:"attr_value_start,omitempty"`
	// End value of interval attribute
	AttrValueEnd string `json:"attr_value_end,omitempty" xml:"attr_value_end,omitempty"`
	// Attribute unit
	AttrValueUnit string `json:"attr_value_unit,omitempty" xml:"attr_value_unit,omitempty"`
	// Attribute ID
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
	// Attribute ID
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
}
