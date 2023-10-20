package aedropshiper

import (
	"sync"
)

// AeopSingleOrderQuery 结构体
type AeopSingleOrderQuery struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolAeopSingleOrderQuery = sync.Pool{
	New: func() any {
		return new(AeopSingleOrderQuery)
	},
}

// GetAeopSingleOrderQuery() 从对象池中获取AeopSingleOrderQuery
func GetAeopSingleOrderQuery() *AeopSingleOrderQuery {
	return poolAeopSingleOrderQuery.Get().(*AeopSingleOrderQuery)
}

// ReleaseAeopSingleOrderQuery 释放AeopSingleOrderQuery
func ReleaseAeopSingleOrderQuery(v *AeopSingleOrderQuery) {
	v.OrderId = 0
	poolAeopSingleOrderQuery.Put(v)
}
