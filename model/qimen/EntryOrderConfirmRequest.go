package qimen

import (
	"sync"
)

// EntryOrderConfirmRequest 结构体
type EntryOrderConfirmRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 分批次入库的最后一次回传
	TotalOrders []TotalOrder `json:"totalOrders,omitempty" xml:"totalOrders>total_order,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenEntryorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolEntryOrderConfirmRequest = sync.Pool{
	New: func() any {
		return new(EntryOrderConfirmRequest)
	},
}

// GetEntryOrderConfirmRequest() 从对象池中获取EntryOrderConfirmRequest
func GetEntryOrderConfirmRequest() *EntryOrderConfirmRequest {
	return poolEntryOrderConfirmRequest.Get().(*EntryOrderConfirmRequest)
}

// ReleaseEntryOrderConfirmRequest 释放EntryOrderConfirmRequest
func ReleaseEntryOrderConfirmRequest(v *EntryOrderConfirmRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.TotalOrders = v.TotalOrders[:0]
	v.EntryOrder = nil
	v.ExtendProps = nil
	poolEntryOrderConfirmRequest.Put(v)
}
