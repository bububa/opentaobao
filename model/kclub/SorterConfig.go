package kclub

// SorterConfig 结构体
type SorterConfig struct {
	// 排序顺序
	Order string `json:"order,omitempty" xml:"order,omitempty"`
	// 排序字段
	Field string `json:"field,omitempty" xml:"field,omitempty"`
}
