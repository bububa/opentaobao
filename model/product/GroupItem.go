package product

import (
	"sync"
)

// GroupItem 结构体
type GroupItem struct {
	// 分组名称
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 分组id
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolGroupItem = sync.Pool{
	New: func() any {
		return new(GroupItem)
	},
}

// GetGroupItem() 从对象池中获取GroupItem
func GetGroupItem() *GroupItem {
	return poolGroupItem.Get().(*GroupItem)
}

// ReleaseGroupItem 释放GroupItem
func ReleaseGroupItem(v *GroupItem) {
	v.Text = ""
	v.Value = 0
	poolGroupItem.Put(v)
}
