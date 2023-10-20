package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceUnbindResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceUnbindResult struct {
	// 系统自动生成
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 系统自动生成
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统自动生成
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceUnbindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceUnbindResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceUnbindResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceUnbindResult
func GetAlibabaAilabsTmallgenieAuthDeviceUnbindResult() *AlibabaAilabsTmallgenieAuthDeviceUnbindResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceUnbindResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceUnbindResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindResult 释放AlibabaAilabsTmallgenieAuthDeviceUnbindResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindResult(v *AlibabaAilabsTmallgenieAuthDeviceUnbindResult) {
	v.Result = ""
	v.Message = ""
	v.Code = 0
	v.Success = false
	poolAlibabaAilabsTmallgenieAuthDeviceUnbindResult.Put(v)
}
