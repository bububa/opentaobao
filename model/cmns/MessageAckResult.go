package cmns

import (
	"sync"
)

// MessageAckResult 结构体
type MessageAckResult struct {
	// 消息回复时间
	AckTime string `json:"ack_time,omitempty" xml:"ack_time,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备id
	Did int64 `json:"did,omitempty" xml:"did,omitempty"`
	// 消息id
	Mid int64 `json:"mid,omitempty" xml:"mid,omitempty"`
}

var poolMessageAckResult = sync.Pool{
	New: func() any {
		return new(MessageAckResult)
	},
}

// GetMessageAckResult() 从对象池中获取MessageAckResult
func GetMessageAckResult() *MessageAckResult {
	return poolMessageAckResult.Get().(*MessageAckResult)
}

// ReleaseMessageAckResult 释放MessageAckResult
func ReleaseMessageAckResult(v *MessageAckResult) {
	v.AckTime = ""
	v.Uuid = ""
	v.Did = 0
	v.Mid = 0
	poolMessageAckResult.Put(v)
}
