package moscm

import (
	"sync"
)

// Paginator 结构体
type Paginator struct {
	// 当前页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页显示多少条记录
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPaginator = sync.Pool{
	New: func() any {
		return new(Paginator)
	},
}

// GetPaginator() 从对象池中获取Paginator
func GetPaginator() *Paginator {
	return poolPaginator.Get().(*Paginator)
}

// ReleasePaginator 释放Paginator
func ReleasePaginator(v *Paginator) {
	v.Page = 0
	v.PageSize = 0
	poolPaginator.Put(v)
}
