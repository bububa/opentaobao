package user

import (
	"sync"
)

// RecommendInfo 结构体
type RecommendInfo struct {
	// 推荐具体数据list
	RecommendGuides []RecommendGuide `json:"recommend_guides,omitempty" xml:"recommend_guides>recommend_guide,omitempty"`
}

var poolRecommendInfo = sync.Pool{
	New: func() any {
		return new(RecommendInfo)
	},
}

// GetRecommendInfo() 从对象池中获取RecommendInfo
func GetRecommendInfo() *RecommendInfo {
	return poolRecommendInfo.Get().(*RecommendInfo)
}

// ReleaseRecommendInfo 释放RecommendInfo
func ReleaseRecommendInfo(v *RecommendInfo) {
	v.RecommendGuides = v.RecommendGuides[:0]
	poolRecommendInfo.Put(v)
}
