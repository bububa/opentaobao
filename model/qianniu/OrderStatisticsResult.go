package qianniu

import (
	"sync"
)

// OrderStatisticsResult 结构体
type OrderStatisticsResult struct {
	// tqdj_order_num
	TqdjOrderNum int64 `json:"tqdj_order_num,omitempty" xml:"tqdj_order_num,omitempty"`
	// take_order_num
	TakeOrderNum int64 `json:"take_order_num,omitempty" xml:"take_order_num,omitempty"`
	// home_order_num
	HomeOrderNum int64 `json:"home_order_num,omitempty" xml:"home_order_num,omitempty"`
	// step_order_num
	StepOrderNum int64 `json:"step_order_num,omitempty" xml:"step_order_num,omitempty"`
}

var poolOrderStatisticsResult = sync.Pool{
	New: func() any {
		return new(OrderStatisticsResult)
	},
}

// GetOrderStatisticsResult() 从对象池中获取OrderStatisticsResult
func GetOrderStatisticsResult() *OrderStatisticsResult {
	return poolOrderStatisticsResult.Get().(*OrderStatisticsResult)
}

// ReleaseOrderStatisticsResult 释放OrderStatisticsResult
func ReleaseOrderStatisticsResult(v *OrderStatisticsResult) {
	v.TqdjOrderNum = 0
	v.TakeOrderNum = 0
	v.HomeOrderNum = 0
	v.StepOrderNum = 0
	poolOrderStatisticsResult.Put(v)
}
