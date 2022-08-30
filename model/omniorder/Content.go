package omniorder

// Content 结构体
type Content struct {
	// 取件码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 取件码的值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
