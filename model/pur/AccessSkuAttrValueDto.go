package pur

// AccessSkuAttrValueDto 结构体
type AccessSkuAttrValueDto struct {
	// 属性名称
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// 属性值名称
	AttrValueName string `json:"attr_value_name,omitempty" xml:"attr_value_name,omitempty"`
}
