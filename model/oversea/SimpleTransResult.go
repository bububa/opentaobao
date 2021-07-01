package oversea

// SimpleTransResult 结构体
type SimpleTransResult struct {
	// translatedText
	TranslatedText string `json:"translated_text,omitempty" xml:"translated_text,omitempty"`
	// statusCode
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
