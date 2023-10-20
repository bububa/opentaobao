package promotion

import (
	"sync"
)

// AlibabaWdkCouponAbandonApiResult 结构体
type AlibabaWdkCouponAbandonApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 操作结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaWdkCouponAbandonApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponAbandonApiResult)
	},
}

// GetAlibabaWdkCouponAbandonApiResult() 从对象池中获取AlibabaWdkCouponAbandonApiResult
func GetAlibabaWdkCouponAbandonApiResult() *AlibabaWdkCouponAbandonApiResult {
	return poolAlibabaWdkCouponAbandonApiResult.Get().(*AlibabaWdkCouponAbandonApiResult)
}

// ReleaseAlibabaWdkCouponAbandonApiResult 释放AlibabaWdkCouponAbandonApiResult
func ReleaseAlibabaWdkCouponAbandonApiResult(v *AlibabaWdkCouponAbandonApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaWdkCouponAbandonApiResult.Put(v)
}
