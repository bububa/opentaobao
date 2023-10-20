package simba

import (
	"sync"
)

// PriceSuggestionDto 结构体
type PriceSuggestionDto struct {
	// 关键词id
	Bidwordid string `json:"bidwordid,omitempty" xml:"bidwordid,omitempty"`
	// 关键词原词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 状态码
	Stat string `json:"stat,omitempty" xml:"stat,omitempty"`
	// 昨日信息
	YesterdayInfo *YesterdayInfo `json:"yesterday_info,omitempty" xml:"yesterday_info,omitempty"`
	// 出价建议
	GuidancePrice *GuidancePrice `json:"guidance_price,omitempty" xml:"guidance_price,omitempty"`
}

var poolPriceSuggestionDto = sync.Pool{
	New: func() any {
		return new(PriceSuggestionDto)
	},
}

// GetPriceSuggestionDto() 从对象池中获取PriceSuggestionDto
func GetPriceSuggestionDto() *PriceSuggestionDto {
	return poolPriceSuggestionDto.Get().(*PriceSuggestionDto)
}

// ReleasePriceSuggestionDto 释放PriceSuggestionDto
func ReleasePriceSuggestionDto(v *PriceSuggestionDto) {
	v.Bidwordid = ""
	v.Word = ""
	v.Stat = ""
	v.YesterdayInfo = nil
	v.GuidancePrice = nil
	poolPriceSuggestionDto.Put(v)
}
