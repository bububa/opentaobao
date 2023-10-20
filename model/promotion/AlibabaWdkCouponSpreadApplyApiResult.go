package promotion

import (
	"sync"
)

// AlibabaWdkCouponSpreadApplyApiResult 结构体
type AlibabaWdkCouponSpreadApplyApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回
	Model *CouponTemplateOperateResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponSpreadApplyApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSpreadApplyApiResult)
	},
}

// GetAlibabaWdkCouponSpreadApplyApiResult() 从对象池中获取AlibabaWdkCouponSpreadApplyApiResult
func GetAlibabaWdkCouponSpreadApplyApiResult() *AlibabaWdkCouponSpreadApplyApiResult {
	return poolAlibabaWdkCouponSpreadApplyApiResult.Get().(*AlibabaWdkCouponSpreadApplyApiResult)
}

// ReleaseAlibabaWdkCouponSpreadApplyApiResult 释放AlibabaWdkCouponSpreadApplyApiResult
func ReleaseAlibabaWdkCouponSpreadApplyApiResult(v *AlibabaWdkCouponSpreadApplyApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkCouponSpreadApplyApiResult.Put(v)
}
