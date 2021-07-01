package wdk

// LanguageInfo 结构体
type LanguageInfo struct {
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 程度
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
}
