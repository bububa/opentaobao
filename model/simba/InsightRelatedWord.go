package simba

import (
	"sync"
)

// InsightRelatedWord 结构体
type InsightRelatedWord struct {
	// 相关度
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 相关词
	RelatedWord string `json:"related_word,omitempty" xml:"related_word,omitempty"`
}

var poolInsightRelatedWord = sync.Pool{
	New: func() any {
		return new(InsightRelatedWord)
	},
}

// GetInsightRelatedWord() 从对象池中获取InsightRelatedWord
func GetInsightRelatedWord() *InsightRelatedWord {
	return poolInsightRelatedWord.Get().(*InsightRelatedWord)
}

// ReleaseInsightRelatedWord 释放InsightRelatedWord
func ReleaseInsightRelatedWord(v *InsightRelatedWord) {
	v.Weight = ""
	v.RelatedWord = ""
	poolInsightRelatedWord.Put(v)
}
