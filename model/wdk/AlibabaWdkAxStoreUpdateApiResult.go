package wdk

import (
	"sync"
)

// AlibabaWdkAxStoreUpdateApiResult 结构体
type AlibabaWdkAxStoreUpdateApiResult struct {
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回对象
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkAxStoreUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkAxStoreUpdateApiResult)
	},
}

// GetAlibabaWdkAxStoreUpdateApiResult() 从对象池中获取AlibabaWdkAxStoreUpdateApiResult
func GetAlibabaWdkAxStoreUpdateApiResult() *AlibabaWdkAxStoreUpdateApiResult {
	return poolAlibabaWdkAxStoreUpdateApiResult.Get().(*AlibabaWdkAxStoreUpdateApiResult)
}

// ReleaseAlibabaWdkAxStoreUpdateApiResult 释放AlibabaWdkAxStoreUpdateApiResult
func ReleaseAlibabaWdkAxStoreUpdateApiResult(v *AlibabaWdkAxStoreUpdateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkAxStoreUpdateApiResult.Put(v)
}
