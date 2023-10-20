package tmallservice

import (
	"sync"
)

// ServiceDefinition 结构体
type ServiceDefinition struct {
	// 业务类型
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 服务类型
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
}

var poolServiceDefinition = sync.Pool{
	New: func() any {
		return new(ServiceDefinition)
	},
}

// GetServiceDefinition() 从对象池中获取ServiceDefinition
func GetServiceDefinition() *ServiceDefinition {
	return poolServiceDefinition.Get().(*ServiceDefinition)
}

// ReleaseServiceDefinition 释放ServiceDefinition
func ReleaseServiceDefinition(v *ServiceDefinition) {
	v.BizCode = ""
	v.ServiceCode = ""
	poolServiceDefinition.Put(v)
}
