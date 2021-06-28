package scbp

// KeywordReportDto 
/* model for simplify = false
type KeywordReportDto struct {

    // 总数目
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 返回数据详情
    
    KeywordEffectList  struct {
        KeywordEffectDto  []KeywordEffectDto `json:"keyword_effect_dto,omitempty"`
    } `json:"keyword_effect_list,omitempty"`
    

}
*/

// KeywordReportDto 
type KeywordReportDto struct {

    // 总数目
    TotalCount   int64 `json:"total_count,omitempty"`

    // 返回数据详情
    KeywordEffectList   []KeywordEffectDto `json:"keyword_effect_list,omitempty"`

}
