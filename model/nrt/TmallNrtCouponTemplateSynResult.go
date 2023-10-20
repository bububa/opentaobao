package nrt

import (
	"sync"
)

// TmallNrtCouponTemplateSynResult 结构体
type TmallNrtCouponTemplateSynResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTmallNrtCouponTemplateSynResult = sync.Pool{
	New: func() any {
		return new(TmallNrtCouponTemplateSynResult)
	},
}

// GetTmallNrtCouponTemplateSynResult() 从对象池中获取TmallNrtCouponTemplateSynResult
func GetTmallNrtCouponTemplateSynResult() *TmallNrtCouponTemplateSynResult {
	return poolTmallNrtCouponTemplateSynResult.Get().(*TmallNrtCouponTemplateSynResult)
}

// ReleaseTmallNrtCouponTemplateSynResult 释放TmallNrtCouponTemplateSynResult
func ReleaseTmallNrtCouponTemplateSynResult(v *TmallNrtCouponTemplateSynResult) {
	v.Message = ""
	v.Code = 0
	poolTmallNrtCouponTemplateSynResult.Put(v)
}
