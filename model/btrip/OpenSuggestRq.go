package btrip

// OpenSuggestRq 
type OpenSuggestRq struct {
    // 第三方企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 搜索关键字
    Keyword   string `json:"keyword,omitempty" xml:"keyword,omitempty"`
    // 0国内机场，2国内机场+临近机场，3国际机场
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 商旅开放平台传2
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
}
