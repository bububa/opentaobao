package qimen

import (
	"sync"
)

// OrderLines 结构体
type OrderLines struct {
	// 订单详情
	OrderLine *OrderLine `json:"orderLine,omitempty" xml:"orderLine,omitempty"`
}

var poolOrderLines = sync.Pool{
	New: func() any {
		return new(OrderLines)
	},
}

// GetOrderLines() 从对象池中获取OrderLines
func GetOrderLines() *OrderLines {
	return poolOrderLines.Get().(*OrderLines)
}

// ReleaseOrderLines 释放OrderLines
func ReleaseOrderLines(v *OrderLines) {
	v.OrderLine = nil
	poolOrderLines.Put(v)
}
