package fenxiao

import (
	"sync"
)

// PaginationResult 结构体
type PaginationResult struct {
	// 仓库信息数组
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPaginationResult = sync.Pool{
	New: func() any {
		return new(PaginationResult)
	},
}

// GetPaginationResult() 从对象池中获取PaginationResult
func GetPaginationResult() *PaginationResult {
	return poolPaginationResult.Get().(*PaginationResult)
}

// ReleasePaginationResult 释放PaginationResult
func ReleasePaginationResult(v *PaginationResult) {
	v.Data = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageNo = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.Success = false
	poolPaginationResult.Put(v)
}
