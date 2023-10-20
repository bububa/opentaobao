package wdk

import (
	"sync"
)

// ContainerDo 结构体
type ContainerDo struct {
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 容器编码
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 生产日期
	ProductionDate string `json:"production_date,omitempty" xml:"production_date,omitempty"`
}

var poolContainerDo = sync.Pool{
	New: func() any {
		return new(ContainerDo)
	},
}

// GetContainerDo() 从对象池中获取ContainerDo
func GetContainerDo() *ContainerDo {
	return poolContainerDo.Get().(*ContainerDo)
}

// ReleaseContainerDo 释放ContainerDo
func ReleaseContainerDo(v *ContainerDo) {
	v.Quantity = ""
	v.ContainerCode = ""
	v.ProductionDate = ""
	poolContainerDo.Put(v)
}
