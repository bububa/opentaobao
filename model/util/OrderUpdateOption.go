package util

import (
	"sync"
)

// OrderUpdateOption 结构体
type OrderUpdateOption struct {
	// 业务识别
	BizName string `json:"biz_name,omitempty" xml:"biz_name,omitempty"`
}

var poolOrderUpdateOption = sync.Pool{
	New: func() any {
		return new(OrderUpdateOption)
	},
}

// GetOrderUpdateOption() 从对象池中获取OrderUpdateOption
func GetOrderUpdateOption() *OrderUpdateOption {
	return poolOrderUpdateOption.Get().(*OrderUpdateOption)
}

// ReleaseOrderUpdateOption 释放OrderUpdateOption
func ReleaseOrderUpdateOption(v *OrderUpdateOption) {
	v.BizName = ""
	poolOrderUpdateOption.Put(v)
}
