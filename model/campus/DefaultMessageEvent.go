package campus

import (
	"sync"
)

// DefaultMessageEvent 结构体
type DefaultMessageEvent struct {
	// 消息ID(唯一ID)
	MsgKey string `json:"msg_key,omitempty" xml:"msg_key,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 消息事件
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 消息类型
	EventType string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	// 消息参数
	JsonData string `json:"json_data,omitempty" xml:"json_data,omitempty"`
	// 资源ID
	ResourceId string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 资源URL
	ResourceUrl string `json:"resource_url,omitempty" xml:"resource_url,omitempty"`
	// 时间戳
	EventTimeMillis int64 `json:"event_time_millis,omitempty" xml:"event_time_millis,omitempty"`
	// 接收人用户ID
	Receiver int64 `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 发送人
	Sender int64 `json:"sender,omitempty" xml:"sender,omitempty"`
}

var poolDefaultMessageEvent = sync.Pool{
	New: func() any {
		return new(DefaultMessageEvent)
	},
}

// GetDefaultMessageEvent() 从对象池中获取DefaultMessageEvent
func GetDefaultMessageEvent() *DefaultMessageEvent {
	return poolDefaultMessageEvent.Get().(*DefaultMessageEvent)
}

// ReleaseDefaultMessageEvent 释放DefaultMessageEvent
func ReleaseDefaultMessageEvent(v *DefaultMessageEvent) {
	v.MsgKey = ""
	v.Source = ""
	v.Action = ""
	v.EventType = ""
	v.JsonData = ""
	v.ResourceId = ""
	v.ResourceUrl = ""
	v.EventTimeMillis = 0
	v.Receiver = 0
	v.Sender = 0
	poolDefaultMessageEvent.Put(v)
}
