package wdk

import (
	"sync"
)

// AlibabaWdkAxStoreCreateApiResult 结构体
type AlibabaWdkAxStoreCreateApiResult struct {
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回对象
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkAxStoreCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkAxStoreCreateApiResult)
	},
}

// GetAlibabaWdkAxStoreCreateApiResult() 从对象池中获取AlibabaWdkAxStoreCreateApiResult
func GetAlibabaWdkAxStoreCreateApiResult() *AlibabaWdkAxStoreCreateApiResult {
	return poolAlibabaWdkAxStoreCreateApiResult.Get().(*AlibabaWdkAxStoreCreateApiResult)
}

// ReleaseAlibabaWdkAxStoreCreateApiResult 释放AlibabaWdkAxStoreCreateApiResult
func ReleaseAlibabaWdkAxStoreCreateApiResult(v *AlibabaWdkAxStoreCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkAxStoreCreateApiResult.Put(v)
}
