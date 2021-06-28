package scbp

// TopKeywordListDTO 
/* model for simplify = false
type TopKeywordListDTO struct {

    // 关键词列表
    
    KeywordList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"keyword_list,omitempty"`
    

}
*/

// TopKeywordListDTO 
type TopKeywordListDTO struct {

    // 关键词列表
    KeywordList   []string `json:"keyword_list,omitempty"`

}
