package trade

import (
	"sync"
)

// AlibabaLstVasTradeflowSaveResult 结构体
type AlibabaLstVasTradeflowSaveResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaLstVasTradeflowSaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstVasTradeflowSaveResult)
	},
}

// GetAlibabaLstVasTradeflowSaveResult() 从对象池中获取AlibabaLstVasTradeflowSaveResult
func GetAlibabaLstVasTradeflowSaveResult() *AlibabaLstVasTradeflowSaveResult {
	return poolAlibabaLstVasTradeflowSaveResult.Get().(*AlibabaLstVasTradeflowSaveResult)
}

// ReleaseAlibabaLstVasTradeflowSaveResult 释放AlibabaLstVasTradeflowSaveResult
func ReleaseAlibabaLstVasTradeflowSaveResult(v *AlibabaLstVasTradeflowSaveResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaLstVasTradeflowSaveResult.Put(v)
}
