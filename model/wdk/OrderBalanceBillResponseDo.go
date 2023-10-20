package wdk

import (
	"sync"
)

// OrderBalanceBillResponseDo 结构体
type OrderBalanceBillResponseDo struct {
	// orderBalanceBillDOList
	OrderBalanceBillList []OrderBalanceBillDo `json:"order_balance_bill_list,omitempty" xml:"order_balance_bill_list>order_balance_bill_do,omitempty"`
	// 是否有下一页0：没有 1：有
	HasNext string `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolOrderBalanceBillResponseDo = sync.Pool{
	New: func() any {
		return new(OrderBalanceBillResponseDo)
	},
}

// GetOrderBalanceBillResponseDo() 从对象池中获取OrderBalanceBillResponseDo
func GetOrderBalanceBillResponseDo() *OrderBalanceBillResponseDo {
	return poolOrderBalanceBillResponseDo.Get().(*OrderBalanceBillResponseDo)
}

// ReleaseOrderBalanceBillResponseDo 释放OrderBalanceBillResponseDo
func ReleaseOrderBalanceBillResponseDo(v *OrderBalanceBillResponseDo) {
	v.OrderBalanceBillList = v.OrderBalanceBillList[:0]
	v.HasNext = ""
	poolOrderBalanceBillResponseDo.Put(v)
}
