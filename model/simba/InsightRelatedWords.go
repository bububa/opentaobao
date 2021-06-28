package simba

// InsightRelatedWords 
/* model for simplify = false
type InsightRelatedWords struct {

    // 原词
    
    Bidword   string `json:"bidword,omitempty"`
    

    // 相关词信息列表
    
    RelatedWordItemsList  struct {
        InsightRelatedWord  []InsightRelatedWord `json:"insight_related_word,omitempty"`
    } `json:"related_word_items_list,omitempty"`
    

}
*/

// InsightRelatedWords 
type InsightRelatedWords struct {

    // 原词
    Bidword   string `json:"bidword,omitempty"`

    // 相关词信息列表
    RelatedWordItemsList   []InsightRelatedWord `json:"related_word_items_list,omitempty"`

}
