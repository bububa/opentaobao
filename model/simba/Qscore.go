package simba

// Qscore 
/* model for simplify = false
type Qscore struct {

    // 词质量得分列表
    
    KeywordQscoreList  struct {
        KeywordQscore  []KeywordQscore `json:"keyword_qscore,omitempty"`
    } `json:"keyword_qscore_list,omitempty"`
    

    // 类目出价质量得分
    
    CatmatchQscore   string `json:"catmatch_qscore,omitempty"`
    

}
*/

// Qscore 
type Qscore struct {

    // 词质量得分列表
    KeywordQscoreList   []KeywordQscore `json:"keyword_qscore_list,omitempty"`

    // 类目出价质量得分
    CatmatchQscore   string `json:"catmatch_qscore,omitempty"`

}
