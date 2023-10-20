package promotion

import (
	"sync"
)

// AlibabaWdkCouponSkuRemoveApiResult 结构体
type AlibabaWdkCouponSkuRemoveApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponSkuRemoveApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSkuRemoveApiResult)
	},
}

// GetAlibabaWdkCouponSkuRemoveApiResult() 从对象池中获取AlibabaWdkCouponSkuRemoveApiResult
func GetAlibabaWdkCouponSkuRemoveApiResult() *AlibabaWdkCouponSkuRemoveApiResult {
	return poolAlibabaWdkCouponSkuRemoveApiResult.Get().(*AlibabaWdkCouponSkuRemoveApiResult)
}

// ReleaseAlibabaWdkCouponSkuRemoveApiResult 释放AlibabaWdkCouponSkuRemoveApiResult
func ReleaseAlibabaWdkCouponSkuRemoveApiResult(v *AlibabaWdkCouponSkuRemoveApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponSkuRemoveApiResult.Put(v)
}
