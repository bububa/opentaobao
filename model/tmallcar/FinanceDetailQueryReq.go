package tmallcar

import (
	"sync"
)

// FinanceDetailQueryReq 结构体
type FinanceDetailQueryReq struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolFinanceDetailQueryReq = sync.Pool{
	New: func() any {
		return new(FinanceDetailQueryReq)
	},
}

// GetFinanceDetailQueryReq() 从对象池中获取FinanceDetailQueryReq
func GetFinanceDetailQueryReq() *FinanceDetailQueryReq {
	return poolFinanceDetailQueryReq.Get().(*FinanceDetailQueryReq)
}

// ReleaseFinanceDetailQueryReq 释放FinanceDetailQueryReq
func ReleaseFinanceDetailQueryReq(v *FinanceDetailQueryReq) {
	v.OrderId = 0
	poolFinanceDetailQueryReq.Put(v)
}
