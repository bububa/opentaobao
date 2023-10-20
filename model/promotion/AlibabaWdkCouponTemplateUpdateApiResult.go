package promotion

import (
	"sync"
)

// AlibabaWdkCouponTemplateUpdateApiResult 结构体
type AlibabaWdkCouponTemplateUpdateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponTemplateUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateUpdateApiResult)
	},
}

// GetAlibabaWdkCouponTemplateUpdateApiResult() 从对象池中获取AlibabaWdkCouponTemplateUpdateApiResult
func GetAlibabaWdkCouponTemplateUpdateApiResult() *AlibabaWdkCouponTemplateUpdateApiResult {
	return poolAlibabaWdkCouponTemplateUpdateApiResult.Get().(*AlibabaWdkCouponTemplateUpdateApiResult)
}

// ReleaseAlibabaWdkCouponTemplateUpdateApiResult 释放AlibabaWdkCouponTemplateUpdateApiResult
func ReleaseAlibabaWdkCouponTemplateUpdateApiResult(v *AlibabaWdkCouponTemplateUpdateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaWdkCouponTemplateUpdateApiResult.Put(v)
}
