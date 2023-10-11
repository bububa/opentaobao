package product

// ChildPropertyKeyValueDto 结构体
type ChildPropertyKeyValueDto struct {
	// 属性键
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
