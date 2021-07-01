package scbp

// KeywordReportDto 
type KeywordReportDto struct {
    // 总数目
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回数据详情
    KeywordEffectList   []KeywordEffectDto `json:"keyword_effect_list,omitempty" xml:"keyword_effect_list>keyword_effect_dto,omitempty"`
}
