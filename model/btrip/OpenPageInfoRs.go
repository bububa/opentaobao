package btrip

import (
	"sync"
)

// OpenPageInfoRs 结构体
type OpenPageInfoRs struct {
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

var poolOpenPageInfoRs = sync.Pool{
	New: func() any {
		return new(OpenPageInfoRs)
	},
}

// GetOpenPageInfoRs() 从对象池中获取OpenPageInfoRs
func GetOpenPageInfoRs() *OpenPageInfoRs {
	return poolOpenPageInfoRs.Get().(*OpenPageInfoRs)
}

// ReleaseOpenPageInfoRs 释放OpenPageInfoRs
func ReleaseOpenPageInfoRs(v *OpenPageInfoRs) {
	v.Page = 0
	v.PageSize = 0
	v.TotalNumber = 0
	poolOpenPageInfoRs.Put(v)
}
