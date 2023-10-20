package iot

import (
	"sync"
)

// AlibabaRetailDevicePayUrlGetResult 结构体
type AlibabaRetailDevicePayUrlGetResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailDevicePayUrlGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDevicePayUrlGetResult)
	},
}

// GetAlibabaRetailDevicePayUrlGetResult() 从对象池中获取AlibabaRetailDevicePayUrlGetResult
func GetAlibabaRetailDevicePayUrlGetResult() *AlibabaRetailDevicePayUrlGetResult {
	return poolAlibabaRetailDevicePayUrlGetResult.Get().(*AlibabaRetailDevicePayUrlGetResult)
}

// ReleaseAlibabaRetailDevicePayUrlGetResult 释放AlibabaRetailDevicePayUrlGetResult
func ReleaseAlibabaRetailDevicePayUrlGetResult(v *AlibabaRetailDevicePayUrlGetResult) {
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Module = ""
	v.Success = false
	poolAlibabaRetailDevicePayUrlGetResult.Put(v)
}
