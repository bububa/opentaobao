package scbp

// KeywordUpdateQuery 
/* model for simplify = false
type KeywordUpdateQuery struct {

    // 更新类型
    
    UpdateType   string `json:"update_type,omitempty"`
    

    // 更新信息
    
    UpdateInfo  *struct {
        KeywordInfo  *KeywordInfo `json:"keyword_info,omitempty"`
    } `json:"update_info,omitempty"`
    

    // 关键词集合
    
    KeywordList  struct {
        KeywordInfo  []KeywordInfo `json:"keyword_info,omitempty"`
    } `json:"keyword_list,omitempty"`
    

}
*/

// KeywordUpdateQuery 
type KeywordUpdateQuery struct {

    // 更新类型
    UpdateType   string `json:"update_type,omitempty"`

    // 更新信息
    UpdateInfo   *KeywordInfo `json:"update_info,omitempty"`

    // 关键词集合
    KeywordList   []KeywordInfo `json:"keyword_list,omitempty"`

}
