package fans

import (
	"sync"
)

// PushMessageParamDo 结构体
type PushMessageParamDo struct {
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 品牌名
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 消息类型
	MessageType string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	// mixnick
	MixNick string `json:"mix_nick,omitempty" xml:"mix_nick,omitempty"`
}

var poolPushMessageParamDo = sync.Pool{
	New: func() any {
		return new(PushMessageParamDo)
	},
}

// GetPushMessageParamDo() 从对象池中获取PushMessageParamDo
func GetPushMessageParamDo() *PushMessageParamDo {
	return poolPushMessageParamDo.Get().(*PushMessageParamDo)
}

// ReleasePushMessageParamDo 释放PushMessageParamDo
func ReleasePushMessageParamDo(v *PushMessageParamDo) {
	v.ActivityId = ""
	v.BrandName = ""
	v.MessageType = ""
	v.MixNick = ""
	poolPushMessageParamDo.Put(v)
}
