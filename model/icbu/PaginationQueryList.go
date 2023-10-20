package icbu

import (
	"sync"
)

// PaginationQueryList 结构体
type PaginationQueryList struct {
	// image_list
	List []PhotobankImageDo `json:"list,omitempty" xml:"list>photobank_image_do,omitempty"`
}

var poolPaginationQueryList = sync.Pool{
	New: func() any {
		return new(PaginationQueryList)
	},
}

// GetPaginationQueryList() 从对象池中获取PaginationQueryList
func GetPaginationQueryList() *PaginationQueryList {
	return poolPaginationQueryList.Get().(*PaginationQueryList)
}

// ReleasePaginationQueryList 释放PaginationQueryList
func ReleasePaginationQueryList(v *PaginationQueryList) {
	v.List = v.List[:0]
	poolPaginationQueryList.Put(v)
}
