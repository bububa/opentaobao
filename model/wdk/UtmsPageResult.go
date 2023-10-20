package wdk

import (
	"sync"
)

// UtmsPageResult 结构体
type UtmsPageResult struct {
	// list
	List []ErpBillDto `json:"list,omitempty" xml:"list>erp_bill_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUtmsPageResult = sync.Pool{
	New: func() any {
		return new(UtmsPageResult)
	},
}

// GetUtmsPageResult() 从对象池中获取UtmsPageResult
func GetUtmsPageResult() *UtmsPageResult {
	return poolUtmsPageResult.Get().(*UtmsPageResult)
}

// ReleaseUtmsPageResult 释放UtmsPageResult
func ReleaseUtmsPageResult(v *UtmsPageResult) {
	v.List = v.List[:0]
	v.Code = ""
	v.Msg = ""
	v.TotalCount = 0
	v.Success = false
	poolUtmsPageResult.Put(v)
}
