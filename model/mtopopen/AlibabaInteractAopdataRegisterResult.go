package mtopopen

import (
	"sync"
)

// AlibabaInteractAopdataRegisterResult 结构体
type AlibabaInteractAopdataRegisterResult struct {
	// xx
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// xxx失败
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// xxx失败
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 跟踪请求使用
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaInteractAopdataRegisterResult = sync.Pool{
	New: func() any {
		return new(AlibabaInteractAopdataRegisterResult)
	},
}

// GetAlibabaInteractAopdataRegisterResult() 从对象池中获取AlibabaInteractAopdataRegisterResult
func GetAlibabaInteractAopdataRegisterResult() *AlibabaInteractAopdataRegisterResult {
	return poolAlibabaInteractAopdataRegisterResult.Get().(*AlibabaInteractAopdataRegisterResult)
}

// ReleaseAlibabaInteractAopdataRegisterResult 释放AlibabaInteractAopdataRegisterResult
func ReleaseAlibabaInteractAopdataRegisterResult(v *AlibabaInteractAopdataRegisterResult) {
	v.Data = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.TraceId = ""
	v.Success = false
	poolAlibabaInteractAopdataRegisterResult.Put(v)
}
