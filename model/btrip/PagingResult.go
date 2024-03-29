package btrip

import (
	"sync"
)

// PagingResult 结构体
type PagingResult struct {
	// 员工数组。
	Items []OpenEmployeeInfo `json:"items,omitempty" xml:"items>open_employee_info,omitempty"`
	// 游标分页下一页游标page_token值，当has_more为true时，会同时返回新的page_token，否则不返回page_token。
	PageToken string `json:"page_token,omitempty" xml:"page_token,omitempty"`
	// 本次请求条件下的数据项总量。
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否还有更多数据项。
	HasMore bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
}

var poolPagingResult = sync.Pool{
	New: func() any {
		return new(PagingResult)
	},
}

// GetPagingResult() 从对象池中获取PagingResult
func GetPagingResult() *PagingResult {
	return poolPagingResult.Get().(*PagingResult)
}

// ReleasePagingResult 释放PagingResult
func ReleasePagingResult(v *PagingResult) {
	v.Items = v.Items[:0]
	v.PageToken = ""
	v.Total = 0
	v.HasMore = false
	poolPagingResult.Put(v)
}
