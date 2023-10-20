package user

import (
	"sync"
)

// NormalMessageDto 结构体
type NormalMessageDto struct {
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 接收者ID
	ReceiverId string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
	// 类型
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 时间戳
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolNormalMessageDto = sync.Pool{
	New: func() any {
		return new(NormalMessageDto)
	},
}

// GetNormalMessageDto() 从对象池中获取NormalMessageDto
func GetNormalMessageDto() *NormalMessageDto {
	return poolNormalMessageDto.Get().(*NormalMessageDto)
}

// ReleaseNormalMessageDto 释放NormalMessageDto
func ReleaseNormalMessageDto(v *NormalMessageDto) {
	v.Content = ""
	v.ReceiverId = ""
	v.ContentType = ""
	v.CreateTime = 0
	poolNormalMessageDto.Put(v)
}
