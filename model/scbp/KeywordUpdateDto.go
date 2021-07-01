package scbp

// KeywordUpdateDto 结构体
type KeywordUpdateDto struct {
	// 要改的价格，单位元
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
}
