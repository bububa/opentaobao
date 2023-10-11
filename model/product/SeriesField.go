package product

// SeriesField 结构体
type SeriesField struct {
	// 文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}
