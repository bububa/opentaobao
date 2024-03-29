package feedflow

import (
	"sync"
)

// TargetDto 结构体
type TargetDto struct {
	// 定向名称
	TargetName string `json:"target_name,omitempty" xml:"target_name,omitempty"`
	// 定向描述
	TargetDesc string `json:"target_desc,omitempty" xml:"target_desc,omitempty"`
	// 定向类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 定向id
	TargetId int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
}

var poolTargetDto = sync.Pool{
	New: func() any {
		return new(TargetDto)
	},
}

// GetTargetDto() 从对象池中获取TargetDto
func GetTargetDto() *TargetDto {
	return poolTargetDto.Get().(*TargetDto)
}

// ReleaseTargetDto 释放TargetDto
func ReleaseTargetDto(v *TargetDto) {
	v.TargetName = ""
	v.TargetDesc = ""
	v.TargetType = ""
	v.TargetId = 0
	poolTargetDto.Put(v)
}
