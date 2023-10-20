package xiamicontent

import (
	"sync"
)

// PagingVo 结构体
type PagingVo struct {
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 总数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolPagingVo = sync.Pool{
	New: func() any {
		return new(PagingVo)
	},
}

// GetPagingVo() 从对象池中获取PagingVo
func GetPagingVo() *PagingVo {
	return poolPagingVo.Get().(*PagingVo)
}

// ReleasePagingVo 释放PagingVo
func ReleasePagingVo(v *PagingVo) {
	v.Pages = 0
	v.Count = 0
	v.PageSize = 0
	v.Page = 0
	poolPagingVo.Put(v)
}
