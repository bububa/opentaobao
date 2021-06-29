package scbp

// KeywordUpdateQuery 
type KeywordUpdateQuery struct {
    // 更新类型
    UpdateType   string `json:"update_type,omitempty" xml:"update_type,omitempty"`
    // 更新信息
    UpdateInfo   *KeywordInfo `json:"update_info,omitempty" xml:"update_info,omitempty"`
    // 关键词集合
    KeywordList   []KeywordInfo `json:"keyword_list,omitempty" xml:"keyword_list>keyword_info,omitempty"`
}
