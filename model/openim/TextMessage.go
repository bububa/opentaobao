package openim

import (
	"sync"
)

// TextMessage 结构体
type TextMessage struct {
	// 发送方userid
	FromId string `json:"from_id,omitempty" xml:"from_id,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 接受方userid
	ToId string `json:"to_id,omitempty" xml:"to_id,omitempty"`
	// 消息时间。UTC时间，精确到秒。必须在一个月内
	Time int64 `json:"time,omitempty" xml:"time,omitempty"`
}

var poolTextMessage = sync.Pool{
	New: func() any {
		return new(TextMessage)
	},
}

// GetTextMessage() 从对象池中获取TextMessage
func GetTextMessage() *TextMessage {
	return poolTextMessage.Get().(*TextMessage)
}

// ReleaseTextMessage 释放TextMessage
func ReleaseTextMessage(v *TextMessage) {
	v.FromId = ""
	v.Message = ""
	v.ToId = ""
	v.Time = 0
	poolTextMessage.Put(v)
}
