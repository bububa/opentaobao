package wdk

import (
	"sync"
)

// AlibabaWdkElemeBillGetApiResult 结构体
type AlibabaWdkElemeBillGetApiResult struct {
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 账单信息
	Model *EleBillBo `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkElemeBillGetApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkElemeBillGetApiResult)
	},
}

// GetAlibabaWdkElemeBillGetApiResult() 从对象池中获取AlibabaWdkElemeBillGetApiResult
func GetAlibabaWdkElemeBillGetApiResult() *AlibabaWdkElemeBillGetApiResult {
	return poolAlibabaWdkElemeBillGetApiResult.Get().(*AlibabaWdkElemeBillGetApiResult)
}

// ReleaseAlibabaWdkElemeBillGetApiResult 释放AlibabaWdkElemeBillGetApiResult
func ReleaseAlibabaWdkElemeBillGetApiResult(v *AlibabaWdkElemeBillGetApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkElemeBillGetApiResult.Put(v)
}
