package aesolution

import (
	"sync"
)

// SingleOrderQuery 结构体
type SingleOrderQuery struct {
	// order ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolSingleOrderQuery = sync.Pool{
	New: func() any {
		return new(SingleOrderQuery)
	},
}

// GetSingleOrderQuery() 从对象池中获取SingleOrderQuery
func GetSingleOrderQuery() *SingleOrderQuery {
	return poolSingleOrderQuery.Get().(*SingleOrderQuery)
}

// ReleaseSingleOrderQuery 释放SingleOrderQuery
func ReleaseSingleOrderQuery(v *SingleOrderQuery) {
	v.OrderId = 0
	poolSingleOrderQuery.Put(v)
}
