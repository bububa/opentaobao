package user

import (
	"sync"
)

// RecommendGuide 结构体
type RecommendGuide struct {
	// 推荐类型
	GuideType string `json:"guide_type,omitempty" xml:"guide_type,omitempty"`
	// 推荐引导语
	GuideUtterance string `json:"guide_utterance,omitempty" xml:"guide_utterance,omitempty"`
}

var poolRecommendGuide = sync.Pool{
	New: func() any {
		return new(RecommendGuide)
	},
}

// GetRecommendGuide() 从对象池中获取RecommendGuide
func GetRecommendGuide() *RecommendGuide {
	return poolRecommendGuide.Get().(*RecommendGuide)
}

// ReleaseRecommendGuide 释放RecommendGuide
func ReleaseRecommendGuide(v *RecommendGuide) {
	v.GuideType = ""
	v.GuideUtterance = ""
	poolRecommendGuide.Put(v)
}
