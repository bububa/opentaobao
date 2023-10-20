package user

import (
	"sync"
)

// DownNormalMessageDto 结构体
type DownNormalMessageDto struct {
	// 消息内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息接受者openId
	ReceiverId string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
	// 消息类型
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 轻店铺appId
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 消息创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolDownNormalMessageDto = sync.Pool{
	New: func() any {
		return new(DownNormalMessageDto)
	},
}

// GetDownNormalMessageDto() 从对象池中获取DownNormalMessageDto
func GetDownNormalMessageDto() *DownNormalMessageDto {
	return poolDownNormalMessageDto.Get().(*DownNormalMessageDto)
}

// ReleaseDownNormalMessageDto 释放DownNormalMessageDto
func ReleaseDownNormalMessageDto(v *DownNormalMessageDto) {
	v.Content = ""
	v.ReceiverId = ""
	v.ContentType = ""
	v.AppId = ""
	v.CreateTime = 0
	poolDownNormalMessageDto.Put(v)
}
