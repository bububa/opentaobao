package moscm

import (
	"sync"
)

// PagedList 结构体
type PagedList struct {
	// 数据结果集
	List []Cspudto `json:"list,omitempty" xml:"list>cspudto,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页获取多少条记录
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

var poolPagedList = sync.Pool{
	New: func() any {
		return new(PagedList)
	},
}

// GetPagedList() 从对象池中获取PagedList
func GetPagedList() *PagedList {
	return poolPagedList.Get().(*PagedList)
}

// ReleasePagedList 释放PagedList
func ReleasePagedList(v *PagedList) {
	v.List = v.List[:0]
	v.CurrentPage = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.TotalPage = 0
	poolPagedList.Put(v)
}
