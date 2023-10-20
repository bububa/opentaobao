package simba

import (
	"sync"
)

// InsightRelatedWords 结构体
type InsightRelatedWords struct {
	// 相关词信息列表
	RelatedWordItemsList []InsightRelatedWord `json:"related_word_items_list,omitempty" xml:"related_word_items_list>insight_related_word,omitempty"`
	// 原词
	Bidword string `json:"bidword,omitempty" xml:"bidword,omitempty"`
}

var poolInsightRelatedWords = sync.Pool{
	New: func() any {
		return new(InsightRelatedWords)
	},
}

// GetInsightRelatedWords() 从对象池中获取InsightRelatedWords
func GetInsightRelatedWords() *InsightRelatedWords {
	return poolInsightRelatedWords.Get().(*InsightRelatedWords)
}

// ReleaseInsightRelatedWords 释放InsightRelatedWords
func ReleaseInsightRelatedWords(v *InsightRelatedWords) {
	v.RelatedWordItemsList = v.RelatedWordItemsList[:0]
	v.Bidword = ""
	poolInsightRelatedWords.Put(v)
}
