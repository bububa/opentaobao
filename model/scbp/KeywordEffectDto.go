package scbp

import (
	"sync"
)

// KeywordEffectDto 结构体
type KeywordEffectDto struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 开始时间
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 曝光
	Impr int64 `json:"impr,omitempty" xml:"impr,omitempty"`
	// 点击
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 消耗
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 推广时长
	OnlineMin int64 `json:"online_min,omitempty" xml:"online_min,omitempty"`
}

var poolKeywordEffectDto = sync.Pool{
	New: func() any {
		return new(KeywordEffectDto)
	},
}

// GetKeywordEffectDto() 从对象池中获取KeywordEffectDto
func GetKeywordEffectDto() *KeywordEffectDto {
	return poolKeywordEffectDto.Get().(*KeywordEffectDto)
}

// ReleaseKeywordEffectDto 释放KeywordEffectDto
func ReleaseKeywordEffectDto(v *KeywordEffectDto) {
	v.Keyword = ""
	v.StatDate = ""
	v.Impr = 0
	v.Click = 0
	v.Cost = 0
	v.OnlineMin = 0
	poolKeywordEffectDto.Put(v)
}
