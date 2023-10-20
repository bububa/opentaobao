package inventory

import (
	"sync"
)

// AlibabaRetailDeviceInventorySyncResult 结构体
type AlibabaRetailDeviceInventorySyncResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailDeviceInventorySyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceInventorySyncResult)
	},
}

// GetAlibabaRetailDeviceInventorySyncResult() 从对象池中获取AlibabaRetailDeviceInventorySyncResult
func GetAlibabaRetailDeviceInventorySyncResult() *AlibabaRetailDeviceInventorySyncResult {
	return poolAlibabaRetailDeviceInventorySyncResult.Get().(*AlibabaRetailDeviceInventorySyncResult)
}

// ReleaseAlibabaRetailDeviceInventorySyncResult 释放AlibabaRetailDeviceInventorySyncResult
func ReleaseAlibabaRetailDeviceInventorySyncResult(v *AlibabaRetailDeviceInventorySyncResult) {
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Success = false
	poolAlibabaRetailDeviceInventorySyncResult.Put(v)
}
