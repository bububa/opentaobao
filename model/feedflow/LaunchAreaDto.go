package feedflow

import (
	"sync"
)

// LaunchAreaDto 结构体
type LaunchAreaDto struct {
	// 地址名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolLaunchAreaDto = sync.Pool{
	New: func() any {
		return new(LaunchAreaDto)
	},
}

// GetLaunchAreaDto() 从对象池中获取LaunchAreaDto
func GetLaunchAreaDto() *LaunchAreaDto {
	return poolLaunchAreaDto.Get().(*LaunchAreaDto)
}

// ReleaseLaunchAreaDto 释放LaunchAreaDto
func ReleaseLaunchAreaDto(v *LaunchAreaDto) {
	v.Name = ""
	v.Code = 0
	poolLaunchAreaDto.Put(v)
}
