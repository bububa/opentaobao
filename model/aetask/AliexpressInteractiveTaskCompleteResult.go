package aetask

import (
	"sync"
)

// AliexpressInteractiveTaskCompleteResult 结构体
type AliexpressInteractiveTaskCompleteResult struct {
	// 结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressInteractiveTaskCompleteResult = sync.Pool{
	New: func() any {
		return new(AliexpressInteractiveTaskCompleteResult)
	},
}

// GetAliexpressInteractiveTaskCompleteResult() 从对象池中获取AliexpressInteractiveTaskCompleteResult
func GetAliexpressInteractiveTaskCompleteResult() *AliexpressInteractiveTaskCompleteResult {
	return poolAliexpressInteractiveTaskCompleteResult.Get().(*AliexpressInteractiveTaskCompleteResult)
}

// ReleaseAliexpressInteractiveTaskCompleteResult 释放AliexpressInteractiveTaskCompleteResult
func ReleaseAliexpressInteractiveTaskCompleteResult(v *AliexpressInteractiveTaskCompleteResult) {
	v.Result = ""
	v.ErrorCode = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAliexpressInteractiveTaskCompleteResult.Put(v)
}
