package alitripmerchant

import (
	"sync"
)

// OrderCountDetailVo 结构体
type OrderCountDetailVo struct {
	// 状态枚举 code
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 状态中文值
	OrderStatusStr string `json:"order_status_str,omitempty" xml:"order_status_str,omitempty"`
	// 对应数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolOrderCountDetailVo = sync.Pool{
	New: func() any {
		return new(OrderCountDetailVo)
	},
}

// GetOrderCountDetailVo() 从对象池中获取OrderCountDetailVo
func GetOrderCountDetailVo() *OrderCountDetailVo {
	return poolOrderCountDetailVo.Get().(*OrderCountDetailVo)
}

// ReleaseOrderCountDetailVo 释放OrderCountDetailVo
func ReleaseOrderCountDetailVo(v *OrderCountDetailVo) {
	v.OrderStatus = ""
	v.OrderStatusStr = ""
	v.Count = 0
	poolOrderCountDetailVo.Put(v)
}
