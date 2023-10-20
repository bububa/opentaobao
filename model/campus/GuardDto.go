package campus

import (
	"sync"
)

// GuardDto 结构体
type GuardDto struct {
	// 门禁点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门禁点ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolGuardDto = sync.Pool{
	New: func() any {
		return new(GuardDto)
	},
}

// GetGuardDto() 从对象池中获取GuardDto
func GetGuardDto() *GuardDto {
	return poolGuardDto.Get().(*GuardDto)
}

// ReleaseGuardDto 释放GuardDto
func ReleaseGuardDto(v *GuardDto) {
	v.Name = ""
	v.Id = 0
	poolGuardDto.Put(v)
}
