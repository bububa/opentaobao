package kclub

import (
	"sync"
)

// Paging 结构体
type Paging struct {
	// 数据
	DataList []KcSearchQuestionDto `json:"data_list,omitempty" xml:"data_list>kc_search_question_dto,omitempty"`
	// 行数
	RowCount int64 `json:"row_count,omitempty" xml:"row_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
}

var poolPaging = sync.Pool{
	New: func() any {
		return new(Paging)
	},
}

// GetPaging() 从对象池中获取Paging
func GetPaging() *Paging {
	return poolPaging.Get().(*Paging)
}

// ReleasePaging 释放Paging
func ReleasePaging(v *Paging) {
	v.DataList = v.DataList[:0]
	v.RowCount = 0
	v.PageSize = 0
	v.PageNo = 0
	poolPaging.Put(v)
}
