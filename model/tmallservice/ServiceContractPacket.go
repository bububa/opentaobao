package tmallservice

import (
	"sync"
)

// ServiceContractPacket 结构体
type ServiceContractPacket struct {
	// 合同类服务列表
	ServiceList []ServiceContractDo `json:"service_list,omitempty" xml:"service_list>service_contract_do,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 服务名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolServiceContractPacket = sync.Pool{
	New: func() any {
		return new(ServiceContractPacket)
	},
}

// GetServiceContractPacket() 从对象池中获取ServiceContractPacket
func GetServiceContractPacket() *ServiceContractPacket {
	return poolServiceContractPacket.Get().(*ServiceContractPacket)
}

// ReleaseServiceContractPacket 释放ServiceContractPacket
func ReleaseServiceContractPacket(v *ServiceContractPacket) {
	v.ServiceList = v.ServiceList[:0]
	v.Desc = ""
	v.Name = ""
	poolServiceContractPacket.Put(v)
}
