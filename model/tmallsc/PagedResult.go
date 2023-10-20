package tmallsc

import (
	"sync"
)

// PagedResult 结构体
type PagedResult struct {
	// 结算明细list
	DataList []Datalist `json:"data_list,omitempty" xml:"data_list>datalist,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPagedResult = sync.Pool{
	New: func() any {
		return new(PagedResult)
	},
}

// GetPagedResult() 从对象池中获取PagedResult
func GetPagedResult() *PagedResult {
	return poolPagedResult.Get().(*PagedResult)
}

// ReleasePagedResult 释放PagedResult
func ReleasePagedResult(v *PagedResult) {
	v.DataList = v.DataList[:0]
	v.TotalCount = 0
	v.PageSize = 0
	poolPagedResult.Put(v)
}
