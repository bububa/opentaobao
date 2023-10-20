package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotDevicePostResult 结构体
type AlibabaAliqinFcIotDevicePostResult struct {
	// 响应结果描述
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 响应code 判断以此判断是否提交成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotDevicePostResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotDevicePostResult)
	},
}

// GetAlibabaAliqinFcIotDevicePostResult() 从对象池中获取AlibabaAliqinFcIotDevicePostResult
func GetAlibabaAliqinFcIotDevicePostResult() *AlibabaAliqinFcIotDevicePostResult {
	return poolAlibabaAliqinFcIotDevicePostResult.Get().(*AlibabaAliqinFcIotDevicePostResult)
}

// ReleaseAlibabaAliqinFcIotDevicePostResult 释放AlibabaAliqinFcIotDevicePostResult
func ReleaseAlibabaAliqinFcIotDevicePostResult(v *AlibabaAliqinFcIotDevicePostResult) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIotDevicePostResult.Put(v)
}
