package icbudropshipping

// ProductSkuNameValue 结构体
type ProductSkuNameValue struct {
	// Attributes name
	AttrNameDesc string `json:"attr_name_desc,omitempty" xml:"attr_name_desc,omitempty"`
	// Attributes value
	AttrValueDesc string `json:"attr_value_desc,omitempty" xml:"attr_value_desc,omitempty"`
	// Attributes url
	AttrValueImage string `json:"attr_value_image,omitempty" xml:"attr_value_image,omitempty"`
	// Attributes name id
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
	// Attributes value id
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
}
