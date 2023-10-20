package logistic

import (
	"sync"
)

// ShippingOrderEvent 结构体
type ShippingOrderEvent struct {
	// 运单状态,0:运单创建,10:配送商接单,20:骑手接单,80:骑手到店,30:骑手取餐,40:已完成,90:配送失败
	ShippingState int64 `json:"shipping_state,omitempty" xml:"shipping_state,omitempty"`
	// 状态变更时间
	OccurredAt int64 `json:"occurred_at,omitempty" xml:"occurred_at,omitempty"`
}

var poolShippingOrderEvent = sync.Pool{
	New: func() any {
		return new(ShippingOrderEvent)
	},
}

// GetShippingOrderEvent() 从对象池中获取ShippingOrderEvent
func GetShippingOrderEvent() *ShippingOrderEvent {
	return poolShippingOrderEvent.Get().(*ShippingOrderEvent)
}

// ReleaseShippingOrderEvent 释放ShippingOrderEvent
func ReleaseShippingOrderEvent(v *ShippingOrderEvent) {
	v.ShippingState = 0
	v.OccurredAt = 0
	poolShippingOrderEvent.Put(v)
}
