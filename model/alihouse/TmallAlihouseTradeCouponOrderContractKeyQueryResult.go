package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponOrderContractKeyQueryResult 结构体
type TmallAlihouseTradeCouponOrderContractKeyQueryResult struct {
	// ossKey
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolTmallAlihouseTradeCouponOrderContractKeyQueryResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderContractKeyQueryResult)
	},
}

// GetTmallAlihouseTradeCouponOrderContractKeyQueryResult() 从对象池中获取TmallAlihouseTradeCouponOrderContractKeyQueryResult
func GetTmallAlihouseTradeCouponOrderContractKeyQueryResult() *TmallAlihouseTradeCouponOrderContractKeyQueryResult {
	return poolTmallAlihouseTradeCouponOrderContractKeyQueryResult.Get().(*TmallAlihouseTradeCouponOrderContractKeyQueryResult)
}

// ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryResult 释放TmallAlihouseTradeCouponOrderContractKeyQueryResult
func ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryResult(v *TmallAlihouseTradeCouponOrderContractKeyQueryResult) {
	v.Data = ""
	v.Code = ""
	v.Msg = ""
	poolTmallAlihouseTradeCouponOrderContractKeyQueryResult.Put(v)
}
