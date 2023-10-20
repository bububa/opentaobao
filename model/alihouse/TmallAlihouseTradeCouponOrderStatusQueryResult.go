package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponOrderStatusQueryResult 结构体
type TmallAlihouseTradeCouponOrderStatusQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data *CouponOrderStatusDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTmallAlihouseTradeCouponOrderStatusQueryResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderStatusQueryResult)
	},
}

// GetTmallAlihouseTradeCouponOrderStatusQueryResult() 从对象池中获取TmallAlihouseTradeCouponOrderStatusQueryResult
func GetTmallAlihouseTradeCouponOrderStatusQueryResult() *TmallAlihouseTradeCouponOrderStatusQueryResult {
	return poolTmallAlihouseTradeCouponOrderStatusQueryResult.Get().(*TmallAlihouseTradeCouponOrderStatusQueryResult)
}

// ReleaseTmallAlihouseTradeCouponOrderStatusQueryResult 释放TmallAlihouseTradeCouponOrderStatusQueryResult
func ReleaseTmallAlihouseTradeCouponOrderStatusQueryResult(v *TmallAlihouseTradeCouponOrderStatusQueryResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	poolTmallAlihouseTradeCouponOrderStatusQueryResult.Put(v)
}
