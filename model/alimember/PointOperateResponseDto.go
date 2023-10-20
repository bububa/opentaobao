package alimember

import (
	"sync"
)

// PointOperateResponseDto 结构体
type PointOperateResponseDto struct {
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}

var poolPointOperateResponseDto = sync.Pool{
	New: func() any {
		return new(PointOperateResponseDto)
	},
}

// GetPointOperateResponseDto() 从对象池中获取PointOperateResponseDto
func GetPointOperateResponseDto() *PointOperateResponseDto {
	return poolPointOperateResponseDto.Get().(*PointOperateResponseDto)
}

// ReleasePointOperateResponseDto 释放PointOperateResponseDto
func ReleasePointOperateResponseDto(v *PointOperateResponseDto) {
	v.Token = ""
	poolPointOperateResponseDto.Put(v)
}
