package product

// SalePropValueSort 结构体
type SalePropValueSort struct {
	// 属性值文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 属性值ID
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}
