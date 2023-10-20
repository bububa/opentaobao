package tmallservice

import (
	"sync"
)

// ServiceTaskPacket 结构体
type ServiceTaskPacket struct {
	// 服务工单DO
	ServiceList []ServiceTaskDo `json:"service_list,omitempty" xml:"service_list>service_task_do,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 服务名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolServiceTaskPacket = sync.Pool{
	New: func() any {
		return new(ServiceTaskPacket)
	},
}

// GetServiceTaskPacket() 从对象池中获取ServiceTaskPacket
func GetServiceTaskPacket() *ServiceTaskPacket {
	return poolServiceTaskPacket.Get().(*ServiceTaskPacket)
}

// ReleaseServiceTaskPacket 释放ServiceTaskPacket
func ReleaseServiceTaskPacket(v *ServiceTaskPacket) {
	v.ServiceList = v.ServiceList[:0]
	v.Desc = ""
	v.Name = ""
	poolServiceTaskPacket.Put(v)
}
