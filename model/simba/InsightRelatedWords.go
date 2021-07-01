package simba

// InsightRelatedWords 结构体
type InsightRelatedWords struct {
	// 原词
	Bidword string `json:"bidword,omitempty" xml:"bidword,omitempty"`
	// 相关词信息列表
	RelatedWordItemsList []InsightRelatedWord `json:"related_word_items_list,omitempty" xml:"related_word_items_list>insight_related_word,omitempty"`
}
