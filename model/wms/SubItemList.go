package wms

import (
	"sync"
)

// SubItemList 结构体
type SubItemList struct {
	// 子货品
	SubItem *SubItem `json:"sub_item,omitempty" xml:"sub_item,omitempty"`
}

var poolSubItemList = sync.Pool{
	New: func() any {
		return new(SubItemList)
	},
}

// GetSubItemList() 从对象池中获取SubItemList
func GetSubItemList() *SubItemList {
	return poolSubItemList.Get().(*SubItemList)
}

// ReleaseSubItemList 释放SubItemList
func ReleaseSubItemList(v *SubItemList) {
	v.SubItem = nil
	poolSubItemList.Put(v)
}
