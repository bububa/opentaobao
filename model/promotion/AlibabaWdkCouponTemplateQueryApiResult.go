package promotion

import (
	"sync"
)

// AlibabaWdkCouponTemplateQueryApiResult 结构体
type AlibabaWdkCouponTemplateQueryApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponTemplateQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateQueryApiResult)
	},
}

// GetAlibabaWdkCouponTemplateQueryApiResult() 从对象池中获取AlibabaWdkCouponTemplateQueryApiResult
func GetAlibabaWdkCouponTemplateQueryApiResult() *AlibabaWdkCouponTemplateQueryApiResult {
	return poolAlibabaWdkCouponTemplateQueryApiResult.Get().(*AlibabaWdkCouponTemplateQueryApiResult)
}

// ReleaseAlibabaWdkCouponTemplateQueryApiResult 释放AlibabaWdkCouponTemplateQueryApiResult
func ReleaseAlibabaWdkCouponTemplateQueryApiResult(v *AlibabaWdkCouponTemplateQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponTemplateQueryApiResult.Put(v)
}
