package cmns

import (
	"sync"
)

// MessageResult 结构体
type MessageResult struct {
	// 消息过期时间
	ExpireTime int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 消息预计发送数
	Total2send int64 `json:"total2send,omitempty" xml:"total2send,omitempty"`
	// 消息审核状态1,3:待审核 2:审核通过，处于发送中 4:审核拒绝 5 已撤回 6 已去重
	Audit int64 `json:"audit,omitempty" xml:"audit,omitempty"`
	// 消息ID
	Mid int64 `json:"mid,omitempty" xml:"mid,omitempty"`
	// 消息达到设备数
	Sentcount int64 `json:"sentcount,omitempty" xml:"sentcount,omitempty"`
}

var poolMessageResult = sync.Pool{
	New: func() any {
		return new(MessageResult)
	},
}

// GetMessageResult() 从对象池中获取MessageResult
func GetMessageResult() *MessageResult {
	return poolMessageResult.Get().(*MessageResult)
}

// ReleaseMessageResult 释放MessageResult
func ReleaseMessageResult(v *MessageResult) {
	v.ExpireTime = 0
	v.Total2send = 0
	v.Audit = 0
	v.Mid = 0
	v.Sentcount = 0
	poolMessageResult.Put(v)
}
