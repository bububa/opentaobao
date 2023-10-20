package promotion

import (
	"sync"
)

// AlibabaWdkCouponTemplateTerminateApiResult 结构体
type AlibabaWdkCouponTemplateTerminateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponTemplateTerminateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateTerminateApiResult)
	},
}

// GetAlibabaWdkCouponTemplateTerminateApiResult() 从对象池中获取AlibabaWdkCouponTemplateTerminateApiResult
func GetAlibabaWdkCouponTemplateTerminateApiResult() *AlibabaWdkCouponTemplateTerminateApiResult {
	return poolAlibabaWdkCouponTemplateTerminateApiResult.Get().(*AlibabaWdkCouponTemplateTerminateApiResult)
}

// ReleaseAlibabaWdkCouponTemplateTerminateApiResult 释放AlibabaWdkCouponTemplateTerminateApiResult
func ReleaseAlibabaWdkCouponTemplateTerminateApiResult(v *AlibabaWdkCouponTemplateTerminateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaWdkCouponTemplateTerminateApiResult.Put(v)
}
