package baichuan

import (
	"sync"
)

// AsoQueryDeviceResult 结构体
type AsoQueryDeviceResult struct {
	// result
	Results []AsoDeviceCheckResult `json:"results,omitempty" xml:"results>aso_device_check_result,omitempty"`
	// errorDetail
	ErrorDetail string `json:"error_detail,omitempty" xml:"error_detail,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAsoQueryDeviceResult = sync.Pool{
	New: func() any {
		return new(AsoQueryDeviceResult)
	},
}

// GetAsoQueryDeviceResult() 从对象池中获取AsoQueryDeviceResult
func GetAsoQueryDeviceResult() *AsoQueryDeviceResult {
	return poolAsoQueryDeviceResult.Get().(*AsoQueryDeviceResult)
}

// ReleaseAsoQueryDeviceResult 释放AsoQueryDeviceResult
func ReleaseAsoQueryDeviceResult(v *AsoQueryDeviceResult) {
	v.Results = v.Results[:0]
	v.ErrorDetail = ""
	v.ErrorCode = ""
	v.Success = false
	poolAsoQueryDeviceResult.Put(v)
}
