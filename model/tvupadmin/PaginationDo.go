package tvupadmin

import (
	"sync"
)

// PaginationDo 结构体
type PaginationDo struct {
	// 内容列表
	List []AdvertScheduleDo `json:"list,omitempty" xml:"list>advert_schedule_do,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 单页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPaginationDo = sync.Pool{
	New: func() any {
		return new(PaginationDo)
	},
}

// GetPaginationDo() 从对象池中获取PaginationDo
func GetPaginationDo() *PaginationDo {
	return poolPaginationDo.Get().(*PaginationDo)
}

// ReleasePaginationDo 释放PaginationDo
func ReleasePaginationDo(v *PaginationDo) {
	v.List = v.List[:0]
	v.Total = 0
	v.PageNo = 0
	v.PageSize = 0
	poolPaginationDo.Put(v)
}
