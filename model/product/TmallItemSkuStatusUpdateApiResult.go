package product

import (
	"sync"
)

// TmallItemSkuStatusUpdateApiResult 结构体
type TmallItemSkuStatusUpdateApiResult struct {
	// 错误码集合，如有
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果信息
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallItemSkuStatusUpdateApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemSkuStatusUpdateApiResult)
	},
}

// GetTmallItemSkuStatusUpdateApiResult() 从对象池中获取TmallItemSkuStatusUpdateApiResult
func GetTmallItemSkuStatusUpdateApiResult() *TmallItemSkuStatusUpdateApiResult {
	return poolTmallItemSkuStatusUpdateApiResult.Get().(*TmallItemSkuStatusUpdateApiResult)
}

// ReleaseTmallItemSkuStatusUpdateApiResult 释放TmallItemSkuStatusUpdateApiResult
func ReleaseTmallItemSkuStatusUpdateApiResult(v *TmallItemSkuStatusUpdateApiResult) {
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Success = false
	poolTmallItemSkuStatusUpdateApiResult.Put(v)
}
