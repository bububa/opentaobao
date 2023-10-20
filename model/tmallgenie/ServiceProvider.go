package tmallgenie

import (
	"sync"
)

// ServiceProvider 结构体
type ServiceProvider struct {
	// 提供商
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片地址
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
}

var poolServiceProvider = sync.Pool{
	New: func() any {
		return new(ServiceProvider)
	},
}

// GetServiceProvider() 从对象池中获取ServiceProvider
func GetServiceProvider() *ServiceProvider {
	return poolServiceProvider.Get().(*ServiceProvider)
}

// ReleaseServiceProvider 释放ServiceProvider
func ReleaseServiceProvider(v *ServiceProvider) {
	v.Name = ""
	v.Icon = ""
	poolServiceProvider.Put(v)
}
