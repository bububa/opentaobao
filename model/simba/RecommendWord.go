package simba

import (
	"sync"
)

// RecommendWord 结构体
type RecommendWord struct {
	// 推荐的关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 搜索量
	Pv string `json:"pv,omitempty" xml:"pv,omitempty"`
	// 平均价格
	AveragePrice string `json:"average_price,omitempty" xml:"average_price,omitempty"`
	// 相关度
	Pertinence string `json:"pertinence,omitempty" xml:"pertinence,omitempty"`
}

var poolRecommendWord = sync.Pool{
	New: func() any {
		return new(RecommendWord)
	},
}

// GetRecommendWord() 从对象池中获取RecommendWord
func GetRecommendWord() *RecommendWord {
	return poolRecommendWord.Get().(*RecommendWord)
}

// ReleaseRecommendWord 释放RecommendWord
func ReleaseRecommendWord(v *RecommendWord) {
	v.Word = ""
	v.Pv = ""
	v.AveragePrice = ""
	v.Pertinence = ""
	poolRecommendWord.Put(v)
}
