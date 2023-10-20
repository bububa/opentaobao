package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponOrderAuditRefundResult 结构体
type TmallAlihouseTradeCouponOrderAuditRefundResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTmallAlihouseTradeCouponOrderAuditRefundResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderAuditRefundResult)
	},
}

// GetTmallAlihouseTradeCouponOrderAuditRefundResult() 从对象池中获取TmallAlihouseTradeCouponOrderAuditRefundResult
func GetTmallAlihouseTradeCouponOrderAuditRefundResult() *TmallAlihouseTradeCouponOrderAuditRefundResult {
	return poolTmallAlihouseTradeCouponOrderAuditRefundResult.Get().(*TmallAlihouseTradeCouponOrderAuditRefundResult)
}

// ReleaseTmallAlihouseTradeCouponOrderAuditRefundResult 释放TmallAlihouseTradeCouponOrderAuditRefundResult
func ReleaseTmallAlihouseTradeCouponOrderAuditRefundResult(v *TmallAlihouseTradeCouponOrderAuditRefundResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = false
	poolTmallAlihouseTradeCouponOrderAuditRefundResult.Put(v)
}
