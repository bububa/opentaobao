package cmns

import (
	"sync"
)

// Pagination 结构体
type Pagination struct {
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数，当totalCount=0 时，totalPageCount 必须设置为 1
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 分页查询页码
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页每页数据集数
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolPagination = sync.Pool{
	New: func() any {
		return new(Pagination)
	},
}

// GetPagination() 从对象池中获取Pagination
func GetPagination() *Pagination {
	return poolPagination.Get().(*Pagination)
}

// ReleasePagination 释放Pagination
func ReleasePagination(v *Pagination) {
	v.TotalCount = 0
	v.TotalPageCount = 0
	v.PageSize = 0
	v.PageIndex = 0
	poolPagination.Put(v)
}
