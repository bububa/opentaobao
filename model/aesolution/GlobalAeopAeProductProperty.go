package aesolution

// GlobalAeopAeProductProperty 结构体
type GlobalAeopAeProductProperty struct {
	// Customized attribute name
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// Attribute Name ID
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
	// Customized attribute value
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	// Attribute Value ID
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
	// unit of customized attribute value
	AttrValueUnit string `json:"attr_value_unit,omitempty" xml:"attr_value_unit,omitempty"`
}
