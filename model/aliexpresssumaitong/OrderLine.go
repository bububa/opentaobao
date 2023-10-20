package aliexpresssumaitong

import (
	"sync"
)

// OrderLine 结构体
type OrderLine struct {
	// 子订单扩展字段
	ExtraParams string `json:"extra_params,omitempty" xml:"extra_params,omitempty"`
	// 子订单号
	OrderLineId int64 `json:"order_line_id,omitempty" xml:"order_line_id,omitempty"`
}

var poolOrderLine = sync.Pool{
	New: func() any {
		return new(OrderLine)
	},
}

// GetOrderLine() 从对象池中获取OrderLine
func GetOrderLine() *OrderLine {
	return poolOrderLine.Get().(*OrderLine)
}

// ReleaseOrderLine 释放OrderLine
func ReleaseOrderLine(v *OrderLine) {
	v.ExtraParams = ""
	v.OrderLineId = 0
	poolOrderLine.Put(v)
}
