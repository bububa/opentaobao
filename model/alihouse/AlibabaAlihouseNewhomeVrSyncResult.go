package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeVrSyncResult 结构体
type AlibabaAlihouseNewhomeVrSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 户型ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeVrSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVrSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeVrSyncResult() 从对象池中获取AlibabaAlihouseNewhomeVrSyncResult
func GetAlibabaAlihouseNewhomeVrSyncResult() *AlibabaAlihouseNewhomeVrSyncResult {
	return poolAlibabaAlihouseNewhomeVrSyncResult.Get().(*AlibabaAlihouseNewhomeVrSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeVrSyncResult 释放AlibabaAlihouseNewhomeVrSyncResult
func ReleaseAlibabaAlihouseNewhomeVrSyncResult(v *AlibabaAlihouseNewhomeVrSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeVrSyncResult.Put(v)
}
