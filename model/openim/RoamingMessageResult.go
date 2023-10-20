package openim

import (
	"sync"
)

// RoamingMessageResult 结构体
type RoamingMessageResult struct {
	// 消息列表
	Messages []RoamingMessage `json:"messages,omitempty" xml:"messages>roaming_message,omitempty"`
	// 下次迭代key
	NextKey string `json:"next_key,omitempty" xml:"next_key,omitempty"`
}

var poolRoamingMessageResult = sync.Pool{
	New: func() any {
		return new(RoamingMessageResult)
	},
}

// GetRoamingMessageResult() 从对象池中获取RoamingMessageResult
func GetRoamingMessageResult() *RoamingMessageResult {
	return poolRoamingMessageResult.Get().(*RoamingMessageResult)
}

// ReleaseRoamingMessageResult 释放RoamingMessageResult
func ReleaseRoamingMessageResult(v *RoamingMessageResult) {
	v.Messages = v.Messages[:0]
	v.NextKey = ""
	poolRoamingMessageResult.Put(v)
}
