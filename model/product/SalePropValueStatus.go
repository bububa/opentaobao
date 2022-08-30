package product

// SalePropValueStatus 结构体
type SalePropValueStatus struct {
	// 属性文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 属性值ID
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 是否新品
	IsNew bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
}
