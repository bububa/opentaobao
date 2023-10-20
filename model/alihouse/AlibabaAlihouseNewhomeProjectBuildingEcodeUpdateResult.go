package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult 结构体
type AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult() 从对象池中获取AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult
func GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult() *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult {
	return poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult.Get().(*AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult 释放AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult
func ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult(v *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult.Put(v)
}
