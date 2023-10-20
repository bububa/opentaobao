package product

import (
	"sync"
)

// TmallItemSkuSortUpdateApiResult 结构体
type TmallItemSkuSortUpdateApiResult struct {
	// 错误信息集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallItemSkuSortUpdateApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemSkuSortUpdateApiResult)
	},
}

// GetTmallItemSkuSortUpdateApiResult() 从对象池中获取TmallItemSkuSortUpdateApiResult
func GetTmallItemSkuSortUpdateApiResult() *TmallItemSkuSortUpdateApiResult {
	return poolTmallItemSkuSortUpdateApiResult.Get().(*TmallItemSkuSortUpdateApiResult)
}

// ReleaseTmallItemSkuSortUpdateApiResult 释放TmallItemSkuSortUpdateApiResult
func ReleaseTmallItemSkuSortUpdateApiResult(v *TmallItemSkuSortUpdateApiResult) {
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Success = false
	poolTmallItemSkuSortUpdateApiResult.Put(v)
}
