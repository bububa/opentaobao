package wdk

import (
	"sync"
)

// AlibabaWdkCouponTemplateQueryumpactidApiResult 结构体
type AlibabaWdkCouponTemplateQueryumpactidApiResult struct {
	// 数据
	Models []CouponRelatedResponse `json:"models,omitempty" xml:"models>coupon_related_response,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// true为成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponTemplateQueryumpactidApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateQueryumpactidApiResult)
	},
}

// GetAlibabaWdkCouponTemplateQueryumpactidApiResult() 从对象池中获取AlibabaWdkCouponTemplateQueryumpactidApiResult
func GetAlibabaWdkCouponTemplateQueryumpactidApiResult() *AlibabaWdkCouponTemplateQueryumpactidApiResult {
	return poolAlibabaWdkCouponTemplateQueryumpactidApiResult.Get().(*AlibabaWdkCouponTemplateQueryumpactidApiResult)
}

// ReleaseAlibabaWdkCouponTemplateQueryumpactidApiResult 释放AlibabaWdkCouponTemplateQueryumpactidApiResult
func ReleaseAlibabaWdkCouponTemplateQueryumpactidApiResult(v *AlibabaWdkCouponTemplateQueryumpactidApiResult) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkCouponTemplateQueryumpactidApiResult.Put(v)
}
