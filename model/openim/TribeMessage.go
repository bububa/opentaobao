package openim

import (
	"sync"
)

// TribeMessage 结构体
type TribeMessage struct {
	// 消息内容节点序列
	Content []MessageItem `json:"content,omitempty" xml:"content>message_item,omitempty"`
	// 发送方
	FromId *User `json:"from_id,omitempty" xml:"from_id,omitempty"`
	// 消息时间。UTC时间
	Time int64 `json:"time,omitempty" xml:"time,omitempty"`
	// 消息类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 消息UUID
	Uuid int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolTribeMessage = sync.Pool{
	New: func() any {
		return new(TribeMessage)
	},
}

// GetTribeMessage() 从对象池中获取TribeMessage
func GetTribeMessage() *TribeMessage {
	return poolTribeMessage.Get().(*TribeMessage)
}

// ReleaseTribeMessage 释放TribeMessage
func ReleaseTribeMessage(v *TribeMessage) {
	v.Content = v.Content[:0]
	v.FromId = nil
	v.Time = 0
	v.Type = 0
	v.Uuid = 0
	poolTribeMessage.Put(v)
}
