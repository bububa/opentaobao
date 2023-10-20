package alihouse

import (
	"sync"
)

// MessageInfoDto 结构体
type MessageInfoDto struct {
	// 子账号
	BrokerId string `json:"broker_id,omitempty" xml:"broker_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 子业务类型
	BizSubType int64 `json:"biz_sub_type,omitempty" xml:"biz_sub_type,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 顾问类型 1-置业顾问 2-经纪人
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolMessageInfoDto = sync.Pool{
	New: func() any {
		return new(MessageInfoDto)
	},
}

// GetMessageInfoDto() 从对象池中获取MessageInfoDto
func GetMessageInfoDto() *MessageInfoDto {
	return poolMessageInfoDto.Get().(*MessageInfoDto)
}

// ReleaseMessageInfoDto 释放MessageInfoDto
func ReleaseMessageInfoDto(v *MessageInfoDto) {
	v.BrokerId = ""
	v.Title = ""
	v.Content = ""
	v.BizSubType = 0
	v.BizType = 0
	v.Type = 0
	poolMessageInfoDto.Put(v)
}
