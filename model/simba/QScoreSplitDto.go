package simba

// QScoreSplitDto 
/* model for simplify = false
type QScoreSplitDto struct {

    // 类目质量得分
    
    CatMatchScore   string `json:"cat_match_score,omitempty"`
    

    // 推广组id
    
    AdgroupId   int64 `json:"adgroup_id,omitempty"`
    

    // 关键词新质量得分列表，包含PC和移动的质量分
    
    WordScoreList  struct {
        Wordscorelist  []Wordscorelist `json:"wordscorelist,omitempty"`
    } `json:"word_score_list,omitempty"`
    

}
*/

// QScoreSplitDto 
type QScoreSplitDto struct {

    // 类目质量得分
    CatMatchScore   string `json:"cat_match_score,omitempty"`

    // 推广组id
    AdgroupId   int64 `json:"adgroup_id,omitempty"`

    // 关键词新质量得分列表，包含PC和移动的质量分
    WordScoreList   []Wordscorelist `json:"word_score_list,omitempty"`

}
