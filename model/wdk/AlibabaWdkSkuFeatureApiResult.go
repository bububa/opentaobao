package wdk

import (
	"sync"
)

// AlibabaWdkSkuFeatureApiResult 结构体
type AlibabaWdkSkuFeatureApiResult struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口调用成功标志，不表示业务是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaWdkSkuFeatureApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuFeatureApiResult)
	},
}

// GetAlibabaWdkSkuFeatureApiResult() 从对象池中获取AlibabaWdkSkuFeatureApiResult
func GetAlibabaWdkSkuFeatureApiResult() *AlibabaWdkSkuFeatureApiResult {
	return poolAlibabaWdkSkuFeatureApiResult.Get().(*AlibabaWdkSkuFeatureApiResult)
}

// ReleaseAlibabaWdkSkuFeatureApiResult 释放AlibabaWdkSkuFeatureApiResult
func ReleaseAlibabaWdkSkuFeatureApiResult(v *AlibabaWdkSkuFeatureApiResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaWdkSkuFeatureApiResult.Put(v)
}
