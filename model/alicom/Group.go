package alicom

import (
	"sync"
)

// Group 结构体
type Group struct {
	// 接口
	InterfaceList []Interface `json:"interface_list,omitempty" xml:"interface_list>interface,omitempty"`
	// spu的结构
	SpuMap string `json:"spu_map,omitempty" xml:"spu_map,omitempty"`
	// 是否主接口
	Main bool `json:"main,omitempty" xml:"main,omitempty"`
}

var poolGroup = sync.Pool{
	New: func() any {
		return new(Group)
	},
}

// GetGroup() 从对象池中获取Group
func GetGroup() *Group {
	return poolGroup.Get().(*Group)
}

// ReleaseGroup 释放Group
func ReleaseGroup(v *Group) {
	v.InterfaceList = v.InterfaceList[:0]
	v.SpuMap = ""
	v.Main = false
	poolGroup.Put(v)
}
