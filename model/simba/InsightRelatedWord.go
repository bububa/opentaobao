package simba

// InsightRelatedWord 结构体
type InsightRelatedWord struct {
	// 相关度
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 相关词
	RelatedWord string `json:"related_word,omitempty" xml:"related_word,omitempty"`
}
