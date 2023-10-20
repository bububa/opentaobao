package scbp

import (
	"sync"
)

// KeywordReportDto 结构体
type KeywordReportDto struct {
	// 返回数据详情
	KeywordEffectList []KeywordEffectDto `json:"keyword_effect_list,omitempty" xml:"keyword_effect_list>keyword_effect_dto,omitempty"`
	// 总数目
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolKeywordReportDto = sync.Pool{
	New: func() any {
		return new(KeywordReportDto)
	},
}

// GetKeywordReportDto() 从对象池中获取KeywordReportDto
func GetKeywordReportDto() *KeywordReportDto {
	return poolKeywordReportDto.Get().(*KeywordReportDto)
}

// ReleaseKeywordReportDto 释放KeywordReportDto
func ReleaseKeywordReportDto(v *KeywordReportDto) {
	v.KeywordEffectList = v.KeywordEffectList[:0]
	v.TotalCount = 0
	poolKeywordReportDto.Put(v)
}
