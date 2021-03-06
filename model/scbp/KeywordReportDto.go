package scbp

// KeywordReportDto 结构体
type KeywordReportDto struct {
	// 返回数据详情
	KeywordEffectList []KeywordEffectDto `json:"keyword_effect_list,omitempty" xml:"keyword_effect_list>keyword_effect_dto,omitempty"`
	// 总数目
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
