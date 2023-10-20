package category

import (
	"sync"
)

// ItemPickPagingResult 结构体
type ItemPickPagingResult struct {
	// result
	Results []CategoryDto `json:"results,omitempty" xml:"results>category_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
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
	v.Success = false
	poolItemPickPagingResult.Put(v)
}
