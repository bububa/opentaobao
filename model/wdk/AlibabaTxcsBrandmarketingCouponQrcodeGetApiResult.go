package wdk

import (
	"sync"
)

// AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult 结构体
type AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回内容
	Model *CouponQrcodeResultDo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult)
	},
}

// GetAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult() 从对象池中获取AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult
func GetAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult() *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult {
	return poolAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult.Get().(*AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult)
}

// ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult 释放AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult
func ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult(v *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaTxcsBrandmarketingCouponQrcodeGetApiResult.Put(v)
}
