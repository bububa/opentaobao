package tmallservice

import (
	"sync"
)

// AlibabaServicecenterFulfiltaskCreateResult 结构体
type AlibabaServicecenterFulfiltaskCreateResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 核销单id
	FulfilTaskId int64 `json:"fulfil_task_id,omitempty" xml:"fulfil_task_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaServicecenterFulfiltaskCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterFulfiltaskCreateResult)
	},
}

// GetAlibabaServicecenterFulfiltaskCreateResult() 从对象池中获取AlibabaServicecenterFulfiltaskCreateResult
func GetAlibabaServicecenterFulfiltaskCreateResult() *AlibabaServicecenterFulfiltaskCreateResult {
	return poolAlibabaServicecenterFulfiltaskCreateResult.Get().(*AlibabaServicecenterFulfiltaskCreateResult)
}

// ReleaseAlibabaServicecenterFulfiltaskCreateResult 释放AlibabaServicecenterFulfiltaskCreateResult
func ReleaseAlibabaServicecenterFulfiltaskCreateResult(v *AlibabaServicecenterFulfiltaskCreateResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FulfilTaskId = 0
	v.Success = false
	poolAlibabaServicecenterFulfiltaskCreateResult.Put(v)
}
