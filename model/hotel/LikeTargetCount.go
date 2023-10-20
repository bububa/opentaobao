package hotel

import (
	"sync"
)

// LikeTargetCount 结构体
type LikeTargetCount struct {
	// count
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// targetId
	TargetId int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// voted
	Voted bool `json:"voted,omitempty" xml:"voted,omitempty"`
}

var poolLikeTargetCount = sync.Pool{
	New: func() any {
		return new(LikeTargetCount)
	},
}

// GetLikeTargetCount() 从对象池中获取LikeTargetCount
func GetLikeTargetCount() *LikeTargetCount {
	return poolLikeTargetCount.Get().(*LikeTargetCount)
}

// ReleaseLikeTargetCount 释放LikeTargetCount
func ReleaseLikeTargetCount(v *LikeTargetCount) {
	v.Count = 0
	v.TargetId = 0
	v.Voted = false
	poolLikeTargetCount.Put(v)
}
