package alihouse

import (
	"sync"
)

// AlibabaAlihouseVrfactoryProductionSyncResult 结构体
type AlibabaAlihouseVrfactoryProductionSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回唯一ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseVrfactoryProductionSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseVrfactoryProductionSyncResult)
	},
}

// GetAlibabaAlihouseVrfactoryProductionSyncResult() 从对象池中获取AlibabaAlihouseVrfactoryProductionSyncResult
func GetAlibabaAlihouseVrfactoryProductionSyncResult() *AlibabaAlihouseVrfactoryProductionSyncResult {
	return poolAlibabaAlihouseVrfactoryProductionSyncResult.Get().(*AlibabaAlihouseVrfactoryProductionSyncResult)
}

// ReleaseAlibabaAlihouseVrfactoryProductionSyncResult 释放AlibabaAlihouseVrfactoryProductionSyncResult
func ReleaseAlibabaAlihouseVrfactoryProductionSyncResult(v *AlibabaAlihouseVrfactoryProductionSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseVrfactoryProductionSyncResult.Put(v)
}
