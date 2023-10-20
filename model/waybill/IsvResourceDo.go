package waybill

import (
	"sync"
)

// IsvResourceDo 结构体
type IsvResourceDo struct {
	// 资源内容（当资源类型为TEMPLATE时，为空）
	ResourceContent string `json:"resource_content,omitempty" xml:"resource_content,omitempty"`
	// 资源名称
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 资源类型
	ResourceType string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 资源url（当资源类型为打印项时，为空）
	ResourceUrl string `json:"resource_url,omitempty" xml:"resource_url,omitempty"`
	// 资源id
	ResourceId int64 `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
}

var poolIsvResourceDo = sync.Pool{
	New: func() any {
		return new(IsvResourceDo)
	},
}

// GetIsvResourceDo() 从对象池中获取IsvResourceDo
func GetIsvResourceDo() *IsvResourceDo {
	return poolIsvResourceDo.Get().(*IsvResourceDo)
}

// ReleaseIsvResourceDo 释放IsvResourceDo
func ReleaseIsvResourceDo(v *IsvResourceDo) {
	v.ResourceContent = ""
	v.ResourceName = ""
	v.ResourceType = ""
	v.ResourceUrl = ""
	v.ResourceId = 0
	poolIsvResourceDo.Put(v)
}
