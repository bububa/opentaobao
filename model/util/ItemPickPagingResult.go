package util

import (
	"sync"
)

// ItemPickPagingResult 结构体
type ItemPickPagingResult struct {
	// 返回数据集合
	Results []CountryDto `json:"results,omitempty" xml:"results>country_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页数据条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolItemPickPagingResult = sync.Pool{
	New: func() any {
		return new(ItemPickPagingResult)
	},
}

// GetItemPickPagingResult() 从对象池中获取ItemPickPagingResult
func GetItemPickPagingResult() *ItemPickPagingResult {
	return poolItemPickPagingResult.Get().(*ItemPickPagingResult)
}

// ReleaseItemPickPagingResult 释放ItemPickPagingResult
func ReleaseItemPickPagingResult(v *ItemPickPagingResult) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageNo = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.Success = false
	poolItemPickPagingResult.Put(v)
}
