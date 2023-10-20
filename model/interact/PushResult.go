package interact

import (
	"sync"
)

// PushResult 结构体
type PushResult struct {
	// 新广播ID
	FeedId string `json:"feed_id,omitempty" xml:"feed_id,omitempty"`
}

var poolPushResult = sync.Pool{
	New: func() any {
		return new(PushResult)
	},
}

// GetPushResult() 从对象池中获取PushResult
func GetPushResult() *PushResult {
	return poolPushResult.Get().(*PushResult)
}

// ReleasePushResult 释放PushResult
func ReleasePushResult(v *PushResult) {
	v.FeedId = ""
	poolPushResult.Put(v)
}
