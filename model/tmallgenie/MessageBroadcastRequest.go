package tmallgenie

import (
	"sync"
)

// MessageBroadcastRequest 结构体
type MessageBroadcastRequest struct {
	// 消息号id
	MessageTemplateId string `json:"message_template_id,omitempty" xml:"message_template_id,omitempty"`
	// 消息模板id
	MessageEntityId string `json:"message_entity_id,omitempty" xml:"message_entity_id,omitempty"`
}

var poolMessageBroadcastRequest = sync.Pool{
	New: func() any {
		return new(MessageBroadcastRequest)
	},
}

// GetMessageBroadcastRequest() 从对象池中获取MessageBroadcastRequest
func GetMessageBroadcastRequest() *MessageBroadcastRequest {
	return poolMessageBroadcastRequest.Get().(*MessageBroadcastRequest)
}

// ReleaseMessageBroadcastRequest 释放MessageBroadcastRequest
func ReleaseMessageBroadcastRequest(v *MessageBroadcastRequest) {
	v.MessageTemplateId = ""
	v.MessageEntityId = ""
	poolMessageBroadcastRequest.Put(v)
}
