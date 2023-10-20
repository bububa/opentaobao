package westcrm

import (
	"sync"
)

// OrderVo 结构体
type OrderVo struct {
	// 账单金额
	OrderPay string `json:"order_pay,omitempty" xml:"order_pay,omitempty"`
	// 账单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 消费时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
}

var poolOrderVo = sync.Pool{
	New: func() any {
		return new(OrderVo)
	},
}

// GetOrderVo() 从对象池中获取OrderVo
func GetOrderVo() *OrderVo {
	return poolOrderVo.Get().(*OrderVo)
}

// ReleaseOrderVo 释放OrderVo
func ReleaseOrderVo(v *OrderVo) {
	v.OrderPay = ""
	v.OrderId = ""
	v.OrderTime = ""
	poolOrderVo.Put(v)
}
