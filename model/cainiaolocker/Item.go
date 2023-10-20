package cainiaolocker

import (
	"sync"
)

// Item 结构体
type Item struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolItem = sync.Pool{
	New: func() any {
		return new(Item)
	},
}

// GetItem() 从对象池中获取Item
func GetItem() *Item {
	return poolItem.Get().(*Item)
}

// ReleaseItem 释放Item
func ReleaseItem(v *Item) {
	v.Name = ""
	v.Count = 0
	poolItem.Put(v)
}
