package xiamiatrist

import (
	"sync"
)

// PagingVo 结构体
type PagingVo struct {
	// 每页展示数
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
	v.PageSize = 0
	v.Page = 0
	poolPagingVo.Put(v)
}
