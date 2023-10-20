package hotel

import (
	"sync"
)

// ScoreInfo 结构体
type ScoreInfo struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// label
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// count
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolScoreInfo = sync.Pool{
	New: func() any {
		return new(ScoreInfo)
	},
}

// GetScoreInfo() 从对象池中获取ScoreInfo
func GetScoreInfo() *ScoreInfo {
	return poolScoreInfo.Get().(*ScoreInfo)
}

// ReleaseScoreInfo 释放ScoreInfo
func ReleaseScoreInfo(v *ScoreInfo) {
	v.Desc = ""
	v.Label = ""
	v.Score = ""
	v.Count = 0
	poolScoreInfo.Put(v)
}
