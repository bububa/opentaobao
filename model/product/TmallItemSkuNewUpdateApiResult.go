package product

import (
	"sync"
)

// TmallItemSkuNewUpdateApiResult 结构体
type TmallItemSkuNewUpdateApiResult struct {
	// 错误码信息集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallItemSkuNewUpdateApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemSkuNewUpdateApiResult)
	},
}

// GetTmallItemSkuNewUpdateApiResult() 从对象池中获取TmallItemSkuNewUpdateApiResult
func GetTmallItemSkuNewUpdateApiResult() *TmallItemSkuNewUpdateApiResult {
	return poolTmallItemSkuNewUpdateApiResult.Get().(*TmallItemSkuNewUpdateApiResult)
}

// ReleaseTmallItemSkuNewUpdateApiResult 释放TmallItemSkuNewUpdateApiResult
func ReleaseTmallItemSkuNewUpdateApiResult(v *TmallItemSkuNewUpdateApiResult) {
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Success = false
	poolTmallItemSkuNewUpdateApiResult.Put(v)
}
