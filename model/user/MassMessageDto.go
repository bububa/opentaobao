package user

import (
	"sync"
)

// MassMessageDto 结构体
type MassMessageDto struct {
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 内容类型
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 任务名称
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolMassMessageDto = sync.Pool{
	New: func() any {
		return new(MassMessageDto)
	},
}

// GetMassMessageDto() 从对象池中获取MassMessageDto
func GetMassMessageDto() *MassMessageDto {
	return poolMassMessageDto.Get().(*MassMessageDto)
}

// ReleaseMassMessageDto 释放MassMessageDto
func ReleaseMassMessageDto(v *MassMessageDto) {
	v.Content = ""
	v.ContentType = ""
	v.TaskName = ""
	v.CreateTime = 0
	poolMassMessageDto.Put(v)
}
