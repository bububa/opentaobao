package xiamiatrist

import (
	"sync"
)

// Paging 结构体
type Paging struct {
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 总数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 每页展示数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolPaging = sync.Pool{
	New: func() any {
		return new(Paging)
	},
}

// GetPaging() 从对象池中获取Paging
func GetPaging() *Paging {
	return poolPaging.Get().(*Paging)
}

// ReleasePaging 释放Paging
func ReleasePaging(v *Paging) {
	v.Pages = 0
	v.Count = 0
	v.PageSize = 0
	v.Page = 0
	poolPaging.Put(v)
}
