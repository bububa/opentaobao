package baichuan

import (
	"sync"
)

// AsoActivateDeviceResult 结构体
type AsoActivateDeviceResult struct {
	// errorDetail
	ErrorDetail string `json:"error_detail,omitempty" xml:"error_detail,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAsoActivateDeviceResult = sync.Pool{
	New: func() any {
		return new(AsoActivateDeviceResult)
	},
}

// GetAsoActivateDeviceResult() 从对象池中获取AsoActivateDeviceResult
func GetAsoActivateDeviceResult() *AsoActivateDeviceResult {
	return poolAsoActivateDeviceResult.Get().(*AsoActivateDeviceResult)
}

// ReleaseAsoActivateDeviceResult 释放AsoActivateDeviceResult
func ReleaseAsoActivateDeviceResult(v *AsoActivateDeviceResult) {
	v.ErrorDetail = ""
	v.ErrorCode = ""
	v.Success = false
	poolAsoActivateDeviceResult.Put(v)
}
