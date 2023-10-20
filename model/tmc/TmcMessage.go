package tmc

import (
	"sync"
)

// TmcMessage 结构体
type TmcMessage struct {
	// 用户的昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 消息详细内容，格式为JSON/XML
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息发布时间
	PubTime string `json:"pub_time,omitempty" xml:"pub_time,omitempty"`
	// 消息发布者的AppKey
	PubAppKey string `json:"pub_app_key,omitempty" xml:"pub_app_key,omitempty"`
	// 消息所属主题
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
	// 消息所属的用户编号
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 消息ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTmcMessage = sync.Pool{
	New: func() any {
		return new(TmcMessage)
	},
}

// GetTmcMessage() 从对象池中获取TmcMessage
func GetTmcMessage() *TmcMessage {
	return poolTmcMessage.Get().(*TmcMessage)
}

// ReleaseTmcMessage 释放TmcMessage
func ReleaseTmcMessage(v *TmcMessage) {
	v.UserNick = ""
	v.Content = ""
	v.PubTime = ""
	v.PubAppKey = ""
	v.Topic = ""
	v.UserId = 0
	v.Id = 0
	poolTmcMessage.Put(v)
}
