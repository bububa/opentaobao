package tmallsc

import (
	"sync"
)

// ServiceContent 结构体
type ServiceContent struct {
	// 服务内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolServiceContent = sync.Pool{
	New: func() any {
		return new(ServiceContent)
	},
}

// GetServiceContent() 从对象池中获取ServiceContent
func GetServiceContent() *ServiceContent {
	return poolServiceContent.Get().(*ServiceContent)
}

// ReleaseServiceContent 释放ServiceContent
func ReleaseServiceContent(v *ServiceContent) {
	v.Content = ""
	poolServiceContent.Put(v)
}
