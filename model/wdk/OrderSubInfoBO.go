package wdk

import (
	"sync"
)

// OrderSubInfoBO 结构体
type OrderSubInfoBO struct {
	// 子业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 外部子单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolOrderSubInfoBO = sync.Pool{
	New: func() any {
		return new(OrderSubInfoBO)
	},
}

// GetOrderSubInfoBO() 从对象池中获取OrderSubInfoBO
func GetOrderSubInfoBO() *OrderSubInfoBO {
	return poolOrderSubInfoBO.Get().(*OrderSubInfoBO)
}

// ReleaseOrderSubInfoBO 释放OrderSubInfoBO
func ReleaseOrderSubInfoBO(v *OrderSubInfoBO) {
	v.BizOrderId = ""
	v.OutOrderId = ""
	poolOrderSubInfoBO.Put(v)
}
