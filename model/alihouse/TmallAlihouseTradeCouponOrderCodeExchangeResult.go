package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponOrderCodeExchangeResult 结构体
type TmallAlihouseTradeCouponOrderCodeExchangeResult struct {
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTmallAlihouseTradeCouponOrderCodeExchangeResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderCodeExchangeResult)
	},
}

// GetTmallAlihouseTradeCouponOrderCodeExchangeResult() 从对象池中获取TmallAlihouseTradeCouponOrderCodeExchangeResult
func GetTmallAlihouseTradeCouponOrderCodeExchangeResult() *TmallAlihouseTradeCouponOrderCodeExchangeResult {
	return poolTmallAlihouseTradeCouponOrderCodeExchangeResult.Get().(*TmallAlihouseTradeCouponOrderCodeExchangeResult)
}

// ReleaseTmallAlihouseTradeCouponOrderCodeExchangeResult 释放TmallAlihouseTradeCouponOrderCodeExchangeResult
func ReleaseTmallAlihouseTradeCouponOrderCodeExchangeResult(v *TmallAlihouseTradeCouponOrderCodeExchangeResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = false
	v.IsSuccess = false
	poolTmallAlihouseTradeCouponOrderCodeExchangeResult.Put(v)
}
