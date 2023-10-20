package idleisv

import (
	"sync"
)

// AppraiseIsvOrderQuery 结构体
type AppraiseIsvOrderQuery struct {
	// 交易订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolAppraiseIsvOrderQuery = sync.Pool{
	New: func() any {
		return new(AppraiseIsvOrderQuery)
	},
}

// GetAppraiseIsvOrderQuery() 从对象池中获取AppraiseIsvOrderQuery
func GetAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
	return poolAppraiseIsvOrderQuery.Get().(*AppraiseIsvOrderQuery)
}

// ReleaseAppraiseIsvOrderQuery 释放AppraiseIsvOrderQuery
func ReleaseAppraiseIsvOrderQuery(v *AppraiseIsvOrderQuery) {
	v.BizOrderId = 0
	poolAppraiseIsvOrderQuery.Put(v)
}
