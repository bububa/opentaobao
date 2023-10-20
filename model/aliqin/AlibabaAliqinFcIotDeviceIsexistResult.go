package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotDeviceIsexistResult 结构体
type AlibabaAliqinFcIotDeviceIsexistResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否存在
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotDeviceIsexistResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotDeviceIsexistResult)
	},
}

// GetAlibabaAliqinFcIotDeviceIsexistResult() 从对象池中获取AlibabaAliqinFcIotDeviceIsexistResult
func GetAlibabaAliqinFcIotDeviceIsexistResult() *AlibabaAliqinFcIotDeviceIsexistResult {
	return poolAlibabaAliqinFcIotDeviceIsexistResult.Get().(*AlibabaAliqinFcIotDeviceIsexistResult)
}

// ReleaseAlibabaAliqinFcIotDeviceIsexistResult 释放AlibabaAliqinFcIotDeviceIsexistResult
func ReleaseAlibabaAliqinFcIotDeviceIsexistResult(v *AlibabaAliqinFcIotDeviceIsexistResult) {
	v.Code = ""
	v.Msg = ""
	v.Model = false
	v.Success = false
	poolAlibabaAliqinFcIotDeviceIsexistResult.Put(v)
}
