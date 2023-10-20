package openim

import (
	"sync"
)

// TribeMessageResult 结构体
type TribeMessageResult struct {
	// 消息列表
	Messages []TribeMessage `json:"messages,omitempty" xml:"messages>tribe_message,omitempty"`
	// 迭代key
	NextKey string `json:"next_key,omitempty" xml:"next_key,omitempty"`
}

var poolTribeMessageResult = sync.Pool{
	New: func() any {
		return new(TribeMessageResult)
	},
}

// GetTribeMessageResult() 从对象池中获取TribeMessageResult
func GetTribeMessageResult() *TribeMessageResult {
	return poolTribeMessageResult.Get().(*TribeMessageResult)
}

// ReleaseTribeMessageResult 释放TribeMessageResult
func ReleaseTribeMessageResult(v *TribeMessageResult) {
	v.Messages = v.Messages[:0]
	v.NextKey = ""
	poolTribeMessageResult.Put(v)
}
