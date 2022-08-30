package idleisv

// IdleItemApiPvPairDo 结构体
type IdleItemApiPvPairDo struct {
	// 属性名文本(长度<=30)
	PropertyText string `json:"property_text,omitempty" xml:"property_text,omitempty"`
	// 属性值文本(长度<=30)
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
}
