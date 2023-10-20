package happytrip

import (
	"sync"
)

// OrderCancelResult 结构体
type OrderCancelResult struct {
	// （非必返回字段）扣费金额，如果传force=false且此订单已被司机抢单，就会返回cost字段（此订单未被取消，有可能产生扣费，需要确认），这时如果要强制取消订单,需要再次请求此接口且传force=true,这时此订单会发生扣款（此订单强制取消）
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
}

var poolOrderCancelResult = sync.Pool{
	New: func() any {
		return new(OrderCancelResult)
	},
}

// GetOrderCancelResult() 从对象池中获取OrderCancelResult
func GetOrderCancelResult() *OrderCancelResult {
	return poolOrderCancelResult.Get().(*OrderCancelResult)
}

// ReleaseOrderCancelResult 释放OrderCancelResult
func ReleaseOrderCancelResult(v *OrderCancelResult) {
	v.Cost = ""
	poolOrderCancelResult.Put(v)
}
