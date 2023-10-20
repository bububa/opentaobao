package idleisv

import (
	"sync"
)

// ItemPageQuery 结构体
type ItemPageQuery struct {
	// 类目集
	CategoryIds []int64 `json:"category_ids,omitempty" xml:"category_ids>int64,omitempty"`
	// 商品状态 0:在线 1售出
	Status []string `json:"status,omitempty" xml:"status>string,omitempty"`
	// 页号
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolItemPageQuery = sync.Pool{
	New: func() any {
		return new(ItemPageQuery)
	},
}

// GetItemPageQuery() 从对象池中获取ItemPageQuery
func GetItemPageQuery() *ItemPageQuery {
	return poolItemPageQuery.Get().(*ItemPageQuery)
}

// ReleaseItemPageQuery 释放ItemPageQuery
func ReleaseItemPageQuery(v *ItemPageQuery) {
	v.CategoryIds = v.CategoryIds[:0]
	v.Status = v.Status[:0]
	v.PageNo = 0
	v.PageSize = 0
	poolItemPageQuery.Put(v)
}
