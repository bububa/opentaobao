package lstvending

import (
	"sync"
)

// SortParam 结构体
type SortParam struct {
	// 排序字段名称：gmt_create创建时间、gmt_modified修改时间、id主键
	SortFieldName string `json:"sort_field_name,omitempty" xml:"sort_field_name,omitempty"`
	// 排序方式：asc、desc
	SortOrder string `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
}

var poolSortParam = sync.Pool{
	New: func() any {
		return new(SortParam)
	},
}

// GetSortParam() 从对象池中获取SortParam
func GetSortParam() *SortParam {
	return poolSortParam.Get().(*SortParam)
}

// ReleaseSortParam 释放SortParam
func ReleaseSortParam(v *SortParam) {
	v.SortFieldName = ""
	v.SortOrder = ""
	poolSortParam.Put(v)
}
