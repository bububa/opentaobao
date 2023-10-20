package promotion

import (
	"sync"
)

// AlibabaWdkCouponTemplateCreateApiResult 结构体
type AlibabaWdkCouponTemplateCreateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponTemplateCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateCreateApiResult)
	},
}

// GetAlibabaWdkCouponTemplateCreateApiResult() 从对象池中获取AlibabaWdkCouponTemplateCreateApiResult
func GetAlibabaWdkCouponTemplateCreateApiResult() *AlibabaWdkCouponTemplateCreateApiResult {
	return poolAlibabaWdkCouponTemplateCreateApiResult.Get().(*AlibabaWdkCouponTemplateCreateApiResult)
}

// ReleaseAlibabaWdkCouponTemplateCreateApiResult 释放AlibabaWdkCouponTemplateCreateApiResult
func ReleaseAlibabaWdkCouponTemplateCreateApiResult(v *AlibabaWdkCouponTemplateCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponTemplateCreateApiResult.Put(v)
}
