package simba

// InsightRelatedWords 
type InsightRelatedWords struct {

    // 原词
    Bidword   string `json:"bidword,omitempty"`

    // 相关词信息列表
    RelatedWordItemsList   []InsightRelatedWord `json:"related_word_items_list,omitempty"`

}
