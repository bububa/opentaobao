package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeCouponSubmitResult 结构体
type AlibabaAlihouseNewhomeCouponSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回对象
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeCouponSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCouponSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeCouponSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeCouponSubmitResult
func GetAlibabaAlihouseNewhomeCouponSubmitResult() *AlibabaAlihouseNewhomeCouponSubmitResult {
	return poolAlibabaAlihouseNewhomeCouponSubmitResult.Get().(*AlibabaAlihouseNewhomeCouponSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeCouponSubmitResult 释放AlibabaAlihouseNewhomeCouponSubmitResult
func ReleaseAlibabaAlihouseNewhomeCouponSubmitResult(v *AlibabaAlihouseNewhomeCouponSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeCouponSubmitResult.Put(v)
}
