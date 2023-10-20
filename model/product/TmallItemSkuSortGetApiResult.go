package product

import (
	"sync"
)

// TmallItemSkuSortGetApiResult 结构体
type TmallItemSkuSortGetApiResult struct {
	// 错误信息集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 属性排序结果信息
	Model []ItemSalePropSort `json:"model,omitempty" xml:"model>item_sale_prop_sort,omitempty"`
	// 执行结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallItemSkuSortGetApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemSkuSortGetApiResult)
	},
}

// GetTmallItemSkuSortGetApiResult() 从对象池中获取TmallItemSkuSortGetApiResult
func GetTmallItemSkuSortGetApiResult() *TmallItemSkuSortGetApiResult {
	return poolTmallItemSkuSortGetApiResult.Get().(*TmallItemSkuSortGetApiResult)
}

// ReleaseTmallItemSkuSortGetApiResult 释放TmallItemSkuSortGetApiResult
func ReleaseTmallItemSkuSortGetApiResult(v *TmallItemSkuSortGetApiResult) {
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Model = v.Model[:0]
	v.Success = ""
	poolTmallItemSkuSortGetApiResult.Put(v)
}
