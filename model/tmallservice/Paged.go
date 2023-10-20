package tmallservice

import (
	"sync"
)

// Paged 结构体
type Paged struct {
	// 工单列表
	DataList []SpServiceOrderDto `json:"data_list,omitempty" xml:"data_list>sp_service_order_dto,omitempty"`
	// 总页数
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolPaged = sync.Pool{
	New: func() any {
		return new(Paged)
	},
}

// GetPaged() 从对象池中获取Paged
func GetPaged() *Paged {
	return poolPaged.Get().(*Paged)
}

// ReleasePaged 释放Paged
func ReleasePaged(v *Paged) {
	v.DataList = v.DataList[:0]
	v.TotalPageCount = 0
	v.PageSize = 0
	v.TotalCount = 0
	poolPaged.Put(v)
}
