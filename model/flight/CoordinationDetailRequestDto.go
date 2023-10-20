package flight

import (
	"sync"
)

// CoordinationDetailRequestDto 结构体
type CoordinationDetailRequestDto struct {
	// 协同单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCoordinationDetailRequestDto = sync.Pool{
	New: func() any {
		return new(CoordinationDetailRequestDto)
	},
}

// GetCoordinationDetailRequestDto() 从对象池中获取CoordinationDetailRequestDto
func GetCoordinationDetailRequestDto() *CoordinationDetailRequestDto {
	return poolCoordinationDetailRequestDto.Get().(*CoordinationDetailRequestDto)
}

// ReleaseCoordinationDetailRequestDto 释放CoordinationDetailRequestDto
func ReleaseCoordinationDetailRequestDto(v *CoordinationDetailRequestDto) {
	v.Id = 0
	poolCoordinationDetailRequestDto.Put(v)
}
