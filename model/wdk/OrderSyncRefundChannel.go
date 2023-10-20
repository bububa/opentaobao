package wdk

import (
	"sync"
)

// OrderSyncRefundChannel 结构体
type OrderSyncRefundChannel struct {
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款渠道
	RefundChannel int64 `json:"refund_channel,omitempty" xml:"refund_channel,omitempty"`
	// 退款单id
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}

var poolOrderSyncRefundChannel = sync.Pool{
	New: func() any {
		return new(OrderSyncRefundChannel)
	},
}

// GetOrderSyncRefundChannel() 从对象池中获取OrderSyncRefundChannel
func GetOrderSyncRefundChannel() *OrderSyncRefundChannel {
	return poolOrderSyncRefundChannel.Get().(*OrderSyncRefundChannel)
}

// ReleaseOrderSyncRefundChannel 释放OrderSyncRefundChannel
func ReleaseOrderSyncRefundChannel(v *OrderSyncRefundChannel) {
	v.RefundAmount = 0
	v.RefundChannel = 0
	v.RefundOrderId = 0
	poolOrderSyncRefundChannel.Put(v)
}
