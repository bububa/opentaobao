package wdk

import (
	"sync"
)

// BasePageQuery 结构体
type BasePageQuery struct {
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前分页，从1开始
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
}

var poolBasePageQuery = sync.Pool{
	New: func() any {
		return new(BasePageQuery)
	},
}

// GetBasePageQuery() 从对象池中获取BasePageQuery
func GetBasePageQuery() *BasePageQuery {
	return poolBasePageQuery.Get().(*BasePageQuery)
}

// ReleaseBasePageQuery 释放BasePageQuery
func ReleaseBasePageQuery(v *BasePageQuery) {
	v.PageSize = 0
	v.Current = 0
	poolBasePageQuery.Put(v)
}
