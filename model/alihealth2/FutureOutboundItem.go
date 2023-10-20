package alihealth2

import (
	"sync"
)

// FutureOutboundItem 结构体
type FutureOutboundItem struct {
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolFutureOutboundItem = sync.Pool{
	New: func() any {
		return new(FutureOutboundItem)
	},
}

// GetFutureOutboundItem() 从对象池中获取FutureOutboundItem
func GetFutureOutboundItem() *FutureOutboundItem {
	return poolFutureOutboundItem.Get().(*FutureOutboundItem)
}

// ReleaseFutureOutboundItem 释放FutureOutboundItem
func ReleaseFutureOutboundItem(v *FutureOutboundItem) {
	v.Quantity = 0
	v.ScItemId = 0
	poolFutureOutboundItem.Put(v)
}
