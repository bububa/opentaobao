package wdk

import (
	"sync"
)

// Container 结构体
type Container struct {
	// 容器类型
	ContainerType string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 容器code
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
}

var poolContainer = sync.Pool{
	New: func() any {
		return new(Container)
	},
}

// GetContainer() 从对象池中获取Container
func GetContainer() *Container {
	return poolContainer.Get().(*Container)
}

// ReleaseContainer 释放Container
func ReleaseContainer(v *Container) {
	v.ContainerType = ""
	v.ContainerCode = ""
	poolContainer.Put(v)
}
