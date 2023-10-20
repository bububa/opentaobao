package btrip

import (
	"sync"
)

// PageInfoRs 结构体
type PageInfoRs struct {
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

var poolPageInfoRs = sync.Pool{
	New: func() any {
		return new(PageInfoRs)
	},
}

// GetPageInfoRs() 从对象池中获取PageInfoRs
func GetPageInfoRs() *PageInfoRs {
	return poolPageInfoRs.Get().(*PageInfoRs)
}

// ReleasePageInfoRs 释放PageInfoRs
func ReleasePageInfoRs(v *PageInfoRs) {
	v.Page = 0
	v.PageSize = 0
	v.TotalNumber = 0
	poolPageInfoRs.Put(v)
}
