package wdk

import (
	"sync"
)

// OrderSyncRefundListResult 结构体
type OrderSyncRefundListResult struct {
	// 退款单具体详情
	Orders []OrderSyncRefundDto `json:"orders,omitempty" xml:"orders>order_sync_refund_dto,omitempty"`
	// 结果code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回的信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 下一页
	NextIndex int64 `json:"next_index,omitempty" xml:"next_index,omitempty"`
	// 总数量
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderSyncRefundListResult = sync.Pool{
	New: func() any {
		return new(OrderSyncRefundListResult)
	},
}

// GetOrderSyncRefundListResult() 从对象池中获取OrderSyncRefundListResult
func GetOrderSyncRefundListResult() *OrderSyncRefundListResult {
	return poolOrderSyncRefundListResult.Get().(*OrderSyncRefundListResult)
}

// ReleaseOrderSyncRefundListResult 释放OrderSyncRefundListResult
func ReleaseOrderSyncRefundListResult(v *OrderSyncRefundListResult) {
	v.Orders = v.Orders[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.NextIndex = 0
	v.TotalNumber = 0
	v.Success = false
	poolOrderSyncRefundListResult.Put(v)
}
