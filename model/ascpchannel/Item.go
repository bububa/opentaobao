package ascpchannel

import (
	"sync"
)

// Item 结构体
type Item struct {
	// 商家编码字段
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 后端货品 ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
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
	v.OuterId = ""
	v.ScItemId = 0
	poolItem.Put(v)
}
