package tbitem

// Feature 结构体
type Feature struct {
	// 属性值
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	// 属性键
	AttrKey string `json:"attr_key,omitempty" xml:"attr_key,omitempty"`
}
