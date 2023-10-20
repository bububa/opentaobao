package moscm

import (
	"sync"
)

// OrderCriteria 结构体
type OrderCriteria struct {
	// 订单号
	OrderNumbers []string `json:"order_numbers,omitempty" xml:"order_numbers>string,omitempty"`
	// 未支付(“UNPAID”),已支付(&#34;PAID&#34;),部分发货(&#34;PARTDISTRIBUTION&#34;),全部发货(&#34;ALLDISTRIBUTION&#34;),取消(&#34;CANCEL&#34;)
	Status []string `json:"status,omitempty" xml:"status>string,omitempty"`
	// 银泰专柜Id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 订单创建时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 订单创建时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 供应商专柜Id
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 订单修改时间范围，开始时间
	ModifyEndDate string `json:"modify_end_date,omitempty" xml:"modify_end_date,omitempty"`
	// 订单修改时间范围，结束时间
	ModifyStartDate string `json:"modify_start_date,omitempty" xml:"modify_start_date,omitempty"`
}

var poolOrderCriteria = sync.Pool{
	New: func() any {
		return new(OrderCriteria)
	},
}

// GetOrderCriteria() 从对象池中获取OrderCriteria
func GetOrderCriteria() *OrderCriteria {
	return poolOrderCriteria.Get().(*OrderCriteria)
}

// ReleaseOrderCriteria 释放OrderCriteria
func ReleaseOrderCriteria(v *OrderCriteria) {
	v.OrderNumbers = v.OrderNumbers[:0]
	v.Status = v.Status[:0]
	v.CounterId = ""
	v.StartDate = ""
	v.EndDate = ""
	v.OutCounterId = ""
	v.ModifyEndDate = ""
	v.ModifyStartDate = ""
	poolOrderCriteria.Put(v)
}
