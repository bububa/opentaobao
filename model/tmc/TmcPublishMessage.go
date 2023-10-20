package tmc

import (
	"sync"
)

// TmcPublishMessage 结构体
type TmcPublishMessage struct {
	// 消息内容的JSON表述，必须按照topic的定义来填充
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 目标分组
	TargetGroup string `json:"target_group,omitempty" xml:"target_group,omitempty"`
	// 消息类型
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
}

var poolTmcPublishMessage = sync.Pool{
	New: func() any {
		return new(TmcPublishMessage)
	},
}

// GetTmcPublishMessage() 从对象池中获取TmcPublishMessage
func GetTmcPublishMessage() *TmcPublishMessage {
	return poolTmcPublishMessage.Get().(*TmcPublishMessage)
}

// ReleaseTmcPublishMessage 释放TmcPublishMessage
func ReleaseTmcPublishMessage(v *TmcPublishMessage) {
	v.Content = ""
	v.TargetGroup = ""
	v.Topic = ""
	poolTmcPublishMessage.Put(v)
}
