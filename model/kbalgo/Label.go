package kbalgo

// Label 结构体
type Label struct {
	// 是否外卖
	LabelDescription string `json:"label_description,omitempty" xml:"label_description,omitempty"`
	// 标签类型
	LabelType string `json:"label_type,omitempty" xml:"label_type,omitempty"`
	// 链接
	Schema *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
	// Delivery
	Delivery *Delivery `json:"delivery,omitempty" xml:"delivery,omitempty"`
}
