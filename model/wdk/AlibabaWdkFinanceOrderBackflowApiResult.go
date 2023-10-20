package wdk

import (
	"sync"
)

// AlibabaWdkFinanceOrderBackflowApiResult 结构体
type AlibabaWdkFinanceOrderBackflowApiResult struct {
	// 调用接口返回错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回的具体结果
	Models *FinanceOrderDetailResponse `json:"models,omitempty" xml:"models,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkFinanceOrderBackflowApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFinanceOrderBackflowApiResult)
	},
}

// GetAlibabaWdkFinanceOrderBackflowApiResult() 从对象池中获取AlibabaWdkFinanceOrderBackflowApiResult
func GetAlibabaWdkFinanceOrderBackflowApiResult() *AlibabaWdkFinanceOrderBackflowApiResult {
	return poolAlibabaWdkFinanceOrderBackflowApiResult.Get().(*AlibabaWdkFinanceOrderBackflowApiResult)
}

// ReleaseAlibabaWdkFinanceOrderBackflowApiResult 释放AlibabaWdkFinanceOrderBackflowApiResult
func ReleaseAlibabaWdkFinanceOrderBackflowApiResult(v *AlibabaWdkFinanceOrderBackflowApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Models = nil
	v.Success = false
	poolAlibabaWdkFinanceOrderBackflowApiResult.Put(v)
}
