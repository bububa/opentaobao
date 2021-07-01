package scbp

// ErrorKeyword 结构体
type ErrorKeyword struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 重复关键词
	RepeatKeyword string `json:"repeat_keyword,omitempty" xml:"repeat_keyword,omitempty"`
}
