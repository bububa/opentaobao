package cainiaoncwl

import (
	"sync"
)

// JhFee 结构体
type JhFee struct {
	// 预留字段，总计金额
	TotalLogisticFee int64 `json:"total_logistic_fee,omitempty" xml:"total_logistic_fee,omitempty"`
}

var poolJhFee = sync.Pool{
	New: func() any {
		return new(JhFee)
	},
}

// GetJhFee() 从对象池中获取JhFee
func GetJhFee() *JhFee {
	return poolJhFee.Get().(*JhFee)
}

// ReleaseJhFee 释放JhFee
func ReleaseJhFee(v *JhFee) {
	v.TotalLogisticFee = 0
	poolJhFee.Put(v)
}
