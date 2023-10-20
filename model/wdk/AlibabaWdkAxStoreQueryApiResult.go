package wdk

import (
	"sync"
)

// AlibabaWdkAxStoreQueryApiResult 结构体
type AlibabaWdkAxStoreQueryApiResult struct {
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 查询接口返回对象
	Model *AxStoreQueryResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkAxStoreQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkAxStoreQueryApiResult)
	},
}

// GetAlibabaWdkAxStoreQueryApiResult() 从对象池中获取AlibabaWdkAxStoreQueryApiResult
func GetAlibabaWdkAxStoreQueryApiResult() *AlibabaWdkAxStoreQueryApiResult {
	return poolAlibabaWdkAxStoreQueryApiResult.Get().(*AlibabaWdkAxStoreQueryApiResult)
}

// ReleaseAlibabaWdkAxStoreQueryApiResult 释放AlibabaWdkAxStoreQueryApiResult
func ReleaseAlibabaWdkAxStoreQueryApiResult(v *AlibabaWdkAxStoreQueryApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkAxStoreQueryApiResult.Put(v)
}
