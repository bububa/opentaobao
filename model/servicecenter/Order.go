package servicecenter

import (
	"sync"
)

// Order 结构体
type Order struct {
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 起始日期（订单未付款前可能为空）
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束日期（订单未付款前可能为空）
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单是否可以排班
	SchedulingState bool `json:"scheduling_state,omitempty" xml:"scheduling_state,omitempty"`
}

var poolOrder = sync.Pool{
	New: func() any {
		return new(Order)
	},
}

// GetOrder() 从对象池中获取Order
func GetOrder() *Order {
	return poolOrder.Get().(*Order)
}

// ReleaseOrder 释放Order
func ReleaseOrder(v *Order) {
	v.SellerNick = ""
	v.StartDate = ""
	v.EndDate = ""
	v.OrderId = 0
	v.SchedulingState = false
	poolOrder.Put(v)
}
