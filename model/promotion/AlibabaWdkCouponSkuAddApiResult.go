package promotion

import (
	"sync"
)

// AlibabaWdkCouponSkuAddApiResult 结构体
type AlibabaWdkCouponSkuAddApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponSkuAddApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSkuAddApiResult)
	},
}

// GetAlibabaWdkCouponSkuAddApiResult() 从对象池中获取AlibabaWdkCouponSkuAddApiResult
func GetAlibabaWdkCouponSkuAddApiResult() *AlibabaWdkCouponSkuAddApiResult {
	return poolAlibabaWdkCouponSkuAddApiResult.Get().(*AlibabaWdkCouponSkuAddApiResult)
}

// ReleaseAlibabaWdkCouponSkuAddApiResult 释放AlibabaWdkCouponSkuAddApiResult
func ReleaseAlibabaWdkCouponSkuAddApiResult(v *AlibabaWdkCouponSkuAddApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponSkuAddApiResult.Put(v)
}
