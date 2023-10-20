package openim

import (
	"sync"
)

// RoamingMessageItem 结构体
type RoamingMessageItem struct {
	// 节点类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolRoamingMessageItem = sync.Pool{
	New: func() any {
		return new(RoamingMessageItem)
	},
}

// GetRoamingMessageItem() 从对象池中获取RoamingMessageItem
func GetRoamingMessageItem() *RoamingMessageItem {
	return poolRoamingMessageItem.Get().(*RoamingMessageItem)
}

// ReleaseRoamingMessageItem 释放RoamingMessageItem
func ReleaseRoamingMessageItem(v *RoamingMessageItem) {
	v.Type = ""
	v.Value = ""
	poolRoamingMessageItem.Put(v)
}
