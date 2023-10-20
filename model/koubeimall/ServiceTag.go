package koubeimall

import (
	"sync"
)

// ServiceTag 结构体
type ServiceTag struct {
	// 图标icon
	TagIcon string `json:"tag_icon,omitempty" xml:"tag_icon,omitempty"`
	// 服务标签code
	TagCode string `json:"tag_code,omitempty" xml:"tag_code,omitempty"`
}

var poolServiceTag = sync.Pool{
	New: func() any {
		return new(ServiceTag)
	},
}

// GetServiceTag() 从对象池中获取ServiceTag
func GetServiceTag() *ServiceTag {
	return poolServiceTag.Get().(*ServiceTag)
}

// ReleaseServiceTag 释放ServiceTag
func ReleaseServiceTag(v *ServiceTag) {
	v.TagIcon = ""
	v.TagCode = ""
	poolServiceTag.Put(v)
}
