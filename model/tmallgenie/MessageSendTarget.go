package tmallgenie

import (
	"sync"
)

// MessageSendTarget 结构体
type MessageSendTarget struct {
	// 消息发送用户标识类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 消息发送用户标识id
	TargetIdentity string `json:"target_identity,omitempty" xml:"target_identity,omitempty"`
}

var poolMessageSendTarget = sync.Pool{
	New: func() any {
		return new(MessageSendTarget)
	},
}

// GetMessageSendTarget() 从对象池中获取MessageSendTarget
func GetMessageSendTarget() *MessageSendTarget {
	return poolMessageSendTarget.Get().(*MessageSendTarget)
}

// ReleaseMessageSendTarget 释放MessageSendTarget
func ReleaseMessageSendTarget(v *MessageSendTarget) {
	v.TargetType = ""
	v.TargetIdentity = ""
	poolMessageSendTarget.Put(v)
}
