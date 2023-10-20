package product

import (
	"sync"
)

// TmallItemSkuStatusGetApiResult 结构体
type TmallItemSkuStatusGetApiResult struct {
	// 错误信息
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 返回内容
	ItemSkuStatus *ItemSkuStatus `json:"item_sku_status,omitempty" xml:"item_sku_status,omitempty"`
}

var poolTmallItemSkuStatusGetApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemSkuStatusGetApiResult)
	},
}

// GetTmallItemSkuStatusGetApiResult() 从对象池中获取TmallItemSkuStatusGetApiResult
func GetTmallItemSkuStatusGetApiResult() *TmallItemSkuStatusGetApiResult {
	return poolTmallItemSkuStatusGetApiResult.Get().(*TmallItemSkuStatusGetApiResult)
}

// ReleaseTmallItemSkuStatusGetApiResult 释放TmallItemSkuStatusGetApiResult
func ReleaseTmallItemSkuStatusGetApiResult(v *TmallItemSkuStatusGetApiResult) {
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Success = ""
	v.ItemSkuStatus = nil
	poolTmallItemSkuStatusGetApiResult.Put(v)
}
