package qimen

import (
	"sync"
)

// CombItem 结构体
type CombItem struct {
	// 奇门仓储字段
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 奇门仓储字段
	Count string `json:"count,omitempty" xml:"count,omitempty"`
}

var poolCombItem = sync.Pool{
	New: func() any {
		return new(CombItem)
	},
}

// GetCombItem() 从对象池中获取CombItem
func GetCombItem() *CombItem {
	return poolCombItem.Get().(*CombItem)
}

// ReleaseCombItem 释放CombItem
func ReleaseCombItem(v *CombItem) {
	v.ItemId = ""
	v.Count = ""
	poolCombItem.Put(v)
}
