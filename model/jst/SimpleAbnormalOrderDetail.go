package jst

import (
	"sync"
)

// SimpleAbnormalOrderDetail 结构体
type SimpleAbnormalOrderDetail struct {
	// 订单ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolSimpleAbnormalOrderDetail = sync.Pool{
	New: func() any {
		return new(SimpleAbnormalOrderDetail)
	},
}

// GetSimpleAbnormalOrderDetail() 从对象池中获取SimpleAbnormalOrderDetail
func GetSimpleAbnormalOrderDetail() *SimpleAbnormalOrderDetail {
	return poolSimpleAbnormalOrderDetail.Get().(*SimpleAbnormalOrderDetail)
}

// ReleaseSimpleAbnormalOrderDetail 释放SimpleAbnormalOrderDetail
func ReleaseSimpleAbnormalOrderDetail(v *SimpleAbnormalOrderDetail) {
	v.Tid = 0
	poolSimpleAbnormalOrderDetail.Put(v)
}
