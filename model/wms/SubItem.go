package wms

import (
	"sync"
)

// SubItem 结构体
type SubItem struct {
	// 子货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 子货品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolSubItem = sync.Pool{
	New: func() any {
		return new(SubItem)
	},
}

// GetSubItem() 从对象池中获取SubItem
func GetSubItem() *SubItem {
	return poolSubItem.Get().(*SubItem)
}

// ReleaseSubItem 释放SubItem
func ReleaseSubItem(v *SubItem) {
	v.ScItemId = 0
	v.Count = 0
	poolSubItem.Put(v)
}
