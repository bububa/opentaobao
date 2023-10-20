package shop

import (
	"sync"
)

// AlibabaShopCouponApplyResult 结构体
type AlibabaShopCouponApplyResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 操作成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaShopCouponApplyResult = sync.Pool{
	New: func() any {
		return new(AlibabaShopCouponApplyResult)
	},
}

// GetAlibabaShopCouponApplyResult() 从对象池中获取AlibabaShopCouponApplyResult
func GetAlibabaShopCouponApplyResult() *AlibabaShopCouponApplyResult {
	return poolAlibabaShopCouponApplyResult.Get().(*AlibabaShopCouponApplyResult)
}

// ReleaseAlibabaShopCouponApplyResult 释放AlibabaShopCouponApplyResult
func ReleaseAlibabaShopCouponApplyResult(v *AlibabaShopCouponApplyResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaShopCouponApplyResult.Put(v)
}
