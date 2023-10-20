package promotion

import (
	"sync"
)

// AlibabaWdkCouponSkuQueryApiResult 结构体
type AlibabaWdkCouponSkuQueryApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponSkuQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSkuQueryApiResult)
	},
}

// GetAlibabaWdkCouponSkuQueryApiResult() 从对象池中获取AlibabaWdkCouponSkuQueryApiResult
func GetAlibabaWdkCouponSkuQueryApiResult() *AlibabaWdkCouponSkuQueryApiResult {
	return poolAlibabaWdkCouponSkuQueryApiResult.Get().(*AlibabaWdkCouponSkuQueryApiResult)
}

// ReleaseAlibabaWdkCouponSkuQueryApiResult 释放AlibabaWdkCouponSkuQueryApiResult
func ReleaseAlibabaWdkCouponSkuQueryApiResult(v *AlibabaWdkCouponSkuQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponSkuQueryApiResult.Put(v)
}
