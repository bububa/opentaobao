package security

// Locale 结构体
type Locale struct {
	// 语言代码(参考ISO-639)
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 国家代码(参考ISO-3166)
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}
