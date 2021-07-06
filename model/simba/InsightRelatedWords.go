package simba

// InsightRelatedWords 结构体
type InsightRelatedWords struct {
	// 相关词信息列表
	RelatedWordItemsList []InsightRelatedWord `json:"related_word_items_list,omitempty" xml:"related_word_items_list>insight_related_word,omitempty"`
	// 原词
	Bidword string `json:"bidword,omitempty" xml:"bidword,omitempty"`
}
