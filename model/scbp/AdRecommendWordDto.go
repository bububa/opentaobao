package scbp

import (
	"sync"
)

// AdRecommendWordDto 结构体
type AdRecommendWordDto struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 归一化词
	NormalWord string `json:"normal_word,omitempty" xml:"normal_word,omitempty"`
	// 建议出价
	SuggestPrice string `json:"suggest_price,omitempty" xml:"suggest_price,omitempty"`
	// 同行平均出价
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 市场底价
	BasePrice string `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 搜索热度
	SearchIndex int64 `json:"search_index,omitempty" xml:"search_index,omitempty"`
	// 购买热度
	BuyIndex int64 `json:"buy_index,omitempty" xml:"buy_index,omitempty"`
	// 星级（推广评分）
	Star int64 `json:"star,omitempty" xml:"star,omitempty"`
	// 最优绑定品id
	BestProductId int64 `json:"best_product_id,omitempty" xml:"best_product_id,omitempty"`
}

var poolAdRecommendWordDto = sync.Pool{
	New: func() any {
		return new(AdRecommendWordDto)
	},
}

// GetAdRecommendWordDto() 从对象池中获取AdRecommendWordDto
func GetAdRecommendWordDto() *AdRecommendWordDto {
	return poolAdRecommendWordDto.Get().(*AdRecommendWordDto)
}

// ReleaseAdRecommendWordDto 释放AdRecommendWordDto
func ReleaseAdRecommendWordDto(v *AdRecommendWordDto) {
	v.Keyword = ""
	v.NormalWord = ""
	v.SuggestPrice = ""
	v.AvgPrice = ""
	v.BasePrice = ""
	v.SearchIndex = 0
	v.BuyIndex = 0
	v.Star = 0
	v.BestProductId = 0
	poolAdRecommendWordDto.Put(v)
}
