package ieagency

import (
	"sync"
)

// BasePageDo 结构体
type BasePageDo struct {
	// 第几页
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页大小
	Pageindex int64 `json:"pageindex,omitempty" xml:"pageindex,omitempty"`
	// 总纪录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolBasePageDo = sync.Pool{
	New: func() any {
		return new(BasePageDo)
	},
}

// GetBasePageDo() 从对象池中获取BasePageDo
func GetBasePageDo() *BasePageDo {
	return poolBasePageDo.Get().(*BasePageDo)
}

// ReleaseBasePageDo 释放BasePageDo
func ReleaseBasePageDo(v *BasePageDo) {
	v.PageSize = 0
	v.Pageindex = 0
	v.TotalCount = 0
	poolBasePageDo.Put(v)
}
