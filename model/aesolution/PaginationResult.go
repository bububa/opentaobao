package aesolution

import (
	"sync"
)

// PaginationResult 结构体
type PaginationResult struct {
	// target list
	TargetList []OrderDto `json:"target_list,omitempty" xml:"target_list>order_dto,omitempty"`
	// error massage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// timeStamp
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// total count(SC order is not include the result）
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// the number of each page
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// current page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// total page
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// success or not
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
	v.TargetList = v.TargetList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.TimeStamp = ""
	v.TotalCount = 0
	v.PageSize = 0
	v.CurrentPage = 0
	v.TotalPage = 0
	v.Success = false
	poolPaginationResult.Put(v)
}
