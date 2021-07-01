package paimai

// Feature 结构体
type Feature struct {
	// 属性键
	AttrKey string `json:"attr_key,omitempty" xml:"attr_key,omitempty"`
	// 属性值
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
}
