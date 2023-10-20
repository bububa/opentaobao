package maitix

import (
	"sync"
)

// ProjectPageParam 结构体
type ProjectPageParam struct {
	// 查询页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页数据大小，可以稍微大一点
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolProjectPageParam = sync.Pool{
	New: func() any {
		return new(ProjectPageParam)
	},
}

// GetProjectPageParam() 从对象池中获取ProjectPageParam
func GetProjectPageParam() *ProjectPageParam {
	return poolProjectPageParam.Get().(*ProjectPageParam)
}

// ReleaseProjectPageParam 释放ProjectPageParam
func ReleaseProjectPageParam(v *ProjectPageParam) {
	v.PageNo = 0
	v.PageSize = 0
	poolProjectPageParam.Put(v)
}
