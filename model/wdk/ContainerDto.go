package wdk

import (
	"sync"
)

// ContainerDto 结构体
type ContainerDto struct {
	// 容器类型
	ContainerType string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 容器code
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
}

var poolContainerDto = sync.Pool{
	New: func() any {
		return new(ContainerDto)
	},
}

// GetContainerDto() 从对象池中获取ContainerDto
func GetContainerDto() *ContainerDto {
	return poolContainerDto.Get().(*ContainerDto)
}

// ReleaseContainerDto 释放ContainerDto
func ReleaseContainerDto(v *ContainerDto) {
	v.ContainerType = ""
	v.ContainerCode = ""
	poolContainerDto.Put(v)
}
