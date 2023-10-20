package tmallservice

import (
	"sync"
)

// PagedResult 结构体
type PagedResult struct {
	// 核销单
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
	// 分页数据
	DataList []BillList `json:"data_list,omitempty" xml:"data_list>bill_list,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 当前页条数
	DataCount int64 `json:"data_count,omitempty" xml:"data_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否空页
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
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
	v.Values = v.Values[:0]
	v.DataList = v.DataList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalCount = 0
	v.PageIndex = 0
	v.PageSize = 0
	v.TotalPageCount = 0
	v.DataCount = 0
	v.Success = false
	v.Empty = false
	poolPagedResult.Put(v)
}
