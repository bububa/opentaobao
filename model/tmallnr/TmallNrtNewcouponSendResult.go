package tmallnr

import (
	"sync"
)

// TmallNrtNewcouponSendResult 结构体
type TmallNrtNewcouponSendResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 发券结果对象
	Model *SendCouponResponse `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTmallNrtNewcouponSendResult = sync.Pool{
	New: func() any {
		return new(TmallNrtNewcouponSendResult)
	},
}

// GetTmallNrtNewcouponSendResult() 从对象池中获取TmallNrtNewcouponSendResult
func GetTmallNrtNewcouponSendResult() *TmallNrtNewcouponSendResult {
	return poolTmallNrtNewcouponSendResult.Get().(*TmallNrtNewcouponSendResult)
}

// ReleaseTmallNrtNewcouponSendResult 释放TmallNrtNewcouponSendResult
func ReleaseTmallNrtNewcouponSendResult(v *TmallNrtNewcouponSendResult) {
	v.Code = ""
	v.Message = ""
	v.Model = nil
	poolTmallNrtNewcouponSendResult.Put(v)
}
