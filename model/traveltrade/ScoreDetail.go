package traveltrade

import (
	"sync"
)

// ScoreDetail 结构体
type ScoreDetail struct {
	// 评价内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 评价分数
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 评价维度
	DimensionName string `json:"dimension_name,omitempty" xml:"dimension_name,omitempty"`
}

var poolScoreDetail = sync.Pool{
	New: func() any {
		return new(ScoreDetail)
	},
}

// GetScoreDetail() 从对象池中获取ScoreDetail
func GetScoreDetail() *ScoreDetail {
	return poolScoreDetail.Get().(*ScoreDetail)
}

// ReleaseScoreDetail 释放ScoreDetail
func ReleaseScoreDetail(v *ScoreDetail) {
	v.Content = ""
	v.Score = ""
	v.DimensionName = ""
	poolScoreDetail.Put(v)
}
