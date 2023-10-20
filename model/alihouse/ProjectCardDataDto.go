package alihouse

import (
	"sync"
)

// ProjectCardDataDto 结构体
type ProjectCardDataDto struct {
	// 楼盘外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 卡片占位符
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolProjectCardDataDto = sync.Pool{
	New: func() any {
		return new(ProjectCardDataDto)
	},
}

// GetProjectCardDataDto() 从对象池中获取ProjectCardDataDto
func GetProjectCardDataDto() *ProjectCardDataDto {
	return poolProjectCardDataDto.Get().(*ProjectCardDataDto)
}

// ReleaseProjectCardDataDto 释放ProjectCardDataDto
func ReleaseProjectCardDataDto(v *ProjectCardDataDto) {
	v.OuterId = ""
	v.Key = ""
	poolProjectCardDataDto.Put(v)
}
