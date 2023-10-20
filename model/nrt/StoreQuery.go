package nrt

import (
	"sync"
)

// StoreQuery 结构体
type StoreQuery struct {
	// 卖场Id或者同城id
	StoreIds []int64 `json:"store_ids,omitempty" xml:"store_ids>int64,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolStoreQuery = sync.Pool{
	New: func() any {
		return new(StoreQuery)
	},
}

// GetStoreQuery() 从对象池中获取StoreQuery
func GetStoreQuery() *StoreQuery {
	return poolStoreQuery.Get().(*StoreQuery)
}

// ReleaseStoreQuery 释放StoreQuery
func ReleaseStoreQuery(v *StoreQuery) {
	v.StoreIds = v.StoreIds[:0]
	v.Type = 0
	poolStoreQuery.Put(v)
}
