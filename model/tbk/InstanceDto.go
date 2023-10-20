package tbk

import (
	"sync"
)

// InstanceDto 结构体
type InstanceDto struct {
	// extra
	Extra *Extra `json:"extra,omitempty" xml:"extra,omitempty"`
}

var poolInstanceDto = sync.Pool{
	New: func() any {
		return new(InstanceDto)
	},
}

// GetInstanceDto() 从对象池中获取InstanceDto
func GetInstanceDto() *InstanceDto {
	return poolInstanceDto.Get().(*InstanceDto)
}

// ReleaseInstanceDto 释放InstanceDto
func ReleaseInstanceDto(v *InstanceDto) {
	v.Extra = nil
	poolInstanceDto.Put(v)
}
