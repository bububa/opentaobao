package wdk

import (
	"sync"
)

// AlibabaWdkReverseRefundResult 结构体
type AlibabaWdkReverseRefundResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 退款单ID
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkReverseRefundResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseRefundResult)
	},
}

// GetAlibabaWdkReverseRefundResult() 从对象池中获取AlibabaWdkReverseRefundResult
func GetAlibabaWdkReverseRefundResult() *AlibabaWdkReverseRefundResult {
	return poolAlibabaWdkReverseRefundResult.Get().(*AlibabaWdkReverseRefundResult)
}

// ReleaseAlibabaWdkReverseRefundResult 释放AlibabaWdkReverseRefundResult
func ReleaseAlibabaWdkReverseRefundResult(v *AlibabaWdkReverseRefundResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = 0
	v.Success = false
	poolAlibabaWdkReverseRefundResult.Put(v)
}
