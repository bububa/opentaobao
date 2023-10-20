package logistic

import (
	"sync"
)

// ServicePackage 结构体
type ServicePackage struct {
	// 服务包code
	ServicePackageCode string `json:"service_package_code,omitempty" xml:"service_package_code,omitempty"`
}

var poolServicePackage = sync.Pool{
	New: func() any {
		return new(ServicePackage)
	},
}

// GetServicePackage() 从对象池中获取ServicePackage
func GetServicePackage() *ServicePackage {
	return poolServicePackage.Get().(*ServicePackage)
}

// ReleaseServicePackage 释放ServicePackage
func ReleaseServicePackage(v *ServicePackage) {
	v.ServicePackageCode = ""
	poolServicePackage.Put(v)
}
