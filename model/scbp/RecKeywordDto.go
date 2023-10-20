package scbp

import (
	"sync"
)

// RecKeywordDto 结构体
type RecKeywordDto struct {
	// 同行平均出价，单位元，一位小数
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 购买竞争度[1-6]
	BuyCount string `json:"buy_count,omitempty" xml:"buy_count,omitempty"`
	// 搜索热度[1-6]
	SearchCount string `json:"search_count,omitempty" xml:"search_count,omitempty"`
	// 客户是否已经将该词添加到外贸直通车后台的出价列表中，取值&#39;Y‘和&#39;N’
	IsAdded string `json:"is_added,omitempty" xml:"is_added,omitempty"`
	// 底价，单位元，一位小数
	BasePrice string `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 关键词的推广评分[0-5]
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
	// 与关键词匹配且处于推广中的产品的数量
	MatchCount int64 `json:"match_count,omitempty" xml:"match_count,omitempty"`
}

var poolRecKeywordDto = sync.Pool{
	New: func() any {
		return new(RecKeywordDto)
	},
}

// GetRecKeywordDto() 从对象池中获取RecKeywordDto
func GetRecKeywordDto() *RecKeywordDto {
	return poolRecKeywordDto.Get().(*RecKeywordDto)
}

// ReleaseRecKeywordDto 释放RecKeywordDto
func ReleaseRecKeywordDto(v *RecKeywordDto) {
	v.AvgPrice = ""
	v.Word = ""
	v.BuyCount = ""
	v.SearchCount = ""
	v.IsAdded = ""
	v.BasePrice = ""
	v.QsStar = 0
	v.MatchCount = 0
	poolRecKeywordDto.Put(v)
}
