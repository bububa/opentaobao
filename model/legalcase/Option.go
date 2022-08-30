package legalcase

// Option 结构体
type Option struct {
	// 标题值
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
