package openim

import (
	"sync"
)

// EsMessageResult 结构体
type EsMessageResult struct {
	// 消息序列
	Messages []EsMessage `json:"messages,omitempty" xml:"messages>es_message,omitempty"`
	// nextkey
	NextKey string `json:"next_key,omitempty" xml:"next_key,omitempty"`
}

var poolEsMessageResult = sync.Pool{
	New: func() any {
		return new(EsMessageResult)
	},
}

// GetEsMessageResult() 从对象池中获取EsMessageResult
func GetEsMessageResult() *EsMessageResult {
	return poolEsMessageResult.Get().(*EsMessageResult)
}

// ReleaseEsMessageResult 释放EsMessageResult
func ReleaseEsMessageResult(v *EsMessageResult) {
	v.Messages = v.Messages[:0]
	v.NextKey = ""
	poolEsMessageResult.Put(v)
}
