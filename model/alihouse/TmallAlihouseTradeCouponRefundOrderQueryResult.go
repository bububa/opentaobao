package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponRefundOrderQueryResult 结构体
type TmallAlihouseTradeCouponRefundOrderQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data *CouponRefundOrderStatusDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTmallAlihouseTradeCouponRefundOrderQueryResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponRefundOrderQueryResult)
	},
}

// GetTmallAlihouseTradeCouponRefundOrderQueryResult() 从对象池中获取TmallAlihouseTradeCouponRefundOrderQueryResult
func GetTmallAlihouseTradeCouponRefundOrderQueryResult() *TmallAlihouseTradeCouponRefundOrderQueryResult {
	return poolTmallAlihouseTradeCouponRefundOrderQueryResult.Get().(*TmallAlihouseTradeCouponRefundOrderQueryResult)
}

// ReleaseTmallAlihouseTradeCouponRefundOrderQueryResult 释放TmallAlihouseTradeCouponRefundOrderQueryResult
func ReleaseTmallAlihouseTradeCouponRefundOrderQueryResult(v *TmallAlihouseTradeCouponRefundOrderQueryResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	poolTmallAlihouseTradeCouponRefundOrderQueryResult.Put(v)
}
