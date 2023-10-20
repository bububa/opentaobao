package alihouse

import (
	"sync"
)

// EntrustSellingQuery 结构体
type EntrustSellingQuery struct {
	// 委托业务主ID
	EntrustSellingId int64 `json:"entrust_selling_id,omitempty" xml:"entrust_selling_id,omitempty"`
}

var poolEntrustSellingQuery = sync.Pool{
	New: func() any {
		return new(EntrustSellingQuery)
	},
}

// GetEntrustSellingQuery() 从对象池中获取EntrustSellingQuery
func GetEntrustSellingQuery() *EntrustSellingQuery {
	return poolEntrustSellingQuery.Get().(*EntrustSellingQuery)
}

// ReleaseEntrustSellingQuery 释放EntrustSellingQuery
func ReleaseEntrustSellingQuery(v *EntrustSellingQuery) {
	v.EntrustSellingId = 0
	poolEntrustSellingQuery.Put(v)
}
