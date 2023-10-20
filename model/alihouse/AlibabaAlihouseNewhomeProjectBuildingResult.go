package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectBuildingResult 结构体
type AlibabaAlihouseNewhomeProjectBuildingResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectBuildingResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectBuildingResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingResult() 从对象池中获取AlibabaAlihouseNewhomeProjectBuildingResult
func GetAlibabaAlihouseNewhomeProjectBuildingResult() *AlibabaAlihouseNewhomeProjectBuildingResult {
	return poolAlibabaAlihouseNewhomeProjectBuildingResult.Get().(*AlibabaAlihouseNewhomeProjectBuildingResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingResult 释放AlibabaAlihouseNewhomeProjectBuildingResult
func ReleaseAlibabaAlihouseNewhomeProjectBuildingResult(v *AlibabaAlihouseNewhomeProjectBuildingResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectBuildingResult.Put(v)
}
