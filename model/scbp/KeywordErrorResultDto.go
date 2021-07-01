package scbp

// KeywordErrorResultDto 结构体
type KeywordErrorResultDto struct {
	// reason
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// keywordId
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
}
