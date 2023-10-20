package alilabs

import (
	"sync"
)

// AlibabaAilabsIotDeviceStatusUpdateResult 结构体
type AlibabaAilabsIotDeviceStatusUpdateResult struct {
	// 附加信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 请求响应码，200代表成功
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 设备状态是否更新成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAlibabaAilabsIotDeviceStatusUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceStatusUpdateResult)
	},
}

// GetAlibabaAilabsIotDeviceStatusUpdateResult() 从对象池中获取AlibabaAilabsIotDeviceStatusUpdateResult
func GetAlibabaAilabsIotDeviceStatusUpdateResult() *AlibabaAilabsIotDeviceStatusUpdateResult {
	return poolAlibabaAilabsIotDeviceStatusUpdateResult.Get().(*AlibabaAilabsIotDeviceStatusUpdateResult)
}

// ReleaseAlibabaAilabsIotDeviceStatusUpdateResult 释放AlibabaAilabsIotDeviceStatusUpdateResult
func ReleaseAlibabaAilabsIotDeviceStatusUpdateResult(v *AlibabaAilabsIotDeviceStatusUpdateResult) {
	v.Message = ""
	v.Exception = ""
	v.StatusCode = 0
	v.Result = false
	poolAlibabaAilabsIotDeviceStatusUpdateResult.Put(v)
}
