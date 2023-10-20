package alimsg

import (
	"sync"
)

// SendTemplateMsgResponse 结构体
type SendTemplateMsgResponse struct {
	// 32位的uuid,用来追踪一条消息
	MessageTraceId string `json:"message_trace_id,omitempty" xml:"message_trace_id,omitempty"`
}

var poolSendTemplateMsgResponse = sync.Pool{
	New: func() any {
		return new(SendTemplateMsgResponse)
	},
}

// GetSendTemplateMsgResponse() 从对象池中获取SendTemplateMsgResponse
func GetSendTemplateMsgResponse() *SendTemplateMsgResponse {
	return poolSendTemplateMsgResponse.Get().(*SendTemplateMsgResponse)
}

// ReleaseSendTemplateMsgResponse 释放SendTemplateMsgResponse
func ReleaseSendTemplateMsgResponse(v *SendTemplateMsgResponse) {
	v.MessageTraceId = ""
	poolSendTemplateMsgResponse.Put(v)
}
