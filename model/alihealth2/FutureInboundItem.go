package alihealth2

import (
	"sync"
)

// FutureInboundItem 结构体
type FutureInboundItem struct {
	// 期货入库数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolFutureInboundItem = sync.Pool{
	New: func() any {
		return new(FutureInboundItem)
	},
}

// GetFutureInboundItem() 从对象池中获取FutureInboundItem
func GetFutureInboundItem() *FutureInboundItem {
	return poolFutureInboundItem.Get().(*FutureInboundItem)
}

// ReleaseFutureInboundItem 释放FutureInboundItem
func ReleaseFutureInboundItem(v *FutureInboundItem) {
	v.Quantity = 0
	v.ScItemId = 0
	poolFutureInboundItem.Put(v)
}
