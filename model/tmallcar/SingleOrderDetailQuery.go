package tmallcar

import (
	"sync"
)

// SingleOrderDetailQuery 结构体
type SingleOrderDetailQuery struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolSingleOrderDetailQuery = sync.Pool{
	New: func() any {
		return new(SingleOrderDetailQuery)
	},
}

// GetSingleOrderDetailQuery() 从对象池中获取SingleOrderDetailQuery
func GetSingleOrderDetailQuery() *SingleOrderDetailQuery {
	return poolSingleOrderDetailQuery.Get().(*SingleOrderDetailQuery)
}

// ReleaseSingleOrderDetailQuery 释放SingleOrderDetailQuery
func ReleaseSingleOrderDetailQuery(v *SingleOrderDetailQuery) {
	v.BuyerNick = ""
	v.OrderId = 0
	poolSingleOrderDetailQuery.Put(v)
}
