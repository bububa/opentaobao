package tmallgenie

import (
	"sync"
)

// AlibabaAilabsIotDeviceListGetResult 结构体
type AlibabaAilabsIotDeviceListGetResult struct {
	// 返回值list
	RetValues []RetValue `json:"ret_values,omitempty" xml:"ret_values>ret_value,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAilabsIotDeviceListGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceListGetResult)
	},
}

// GetAlibabaAilabsIotDeviceListGetResult() 从对象池中获取AlibabaAilabsIotDeviceListGetResult
func GetAlibabaAilabsIotDeviceListGetResult() *AlibabaAilabsIotDeviceListGetResult {
	return poolAlibabaAilabsIotDeviceListGetResult.Get().(*AlibabaAilabsIotDeviceListGetResult)
}

// ReleaseAlibabaAilabsIotDeviceListGetResult 释放AlibabaAilabsIotDeviceListGetResult
func ReleaseAlibabaAilabsIotDeviceListGetResult(v *AlibabaAilabsIotDeviceListGetResult) {
	v.RetValues = v.RetValues[:0]
	v.Message = ""
	v.RetCode = 0
	v.Success = false
	poolAlibabaAilabsIotDeviceListGetResult.Put(v)
}
