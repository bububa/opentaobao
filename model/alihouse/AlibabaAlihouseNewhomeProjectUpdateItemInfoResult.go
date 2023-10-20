package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectUpdateItemInfoResult 结构体
type AlibabaAlihouseNewhomeProjectUpdateItemInfoResult struct {
	// 成功描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 楼盘表的主键ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 成功返回值
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectUpdateItemInfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectUpdateItemInfoResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectUpdateItemInfoResult() 从对象池中获取AlibabaAlihouseNewhomeProjectUpdateItemInfoResult
func GetAlibabaAlihouseNewhomeProjectUpdateItemInfoResult() *AlibabaAlihouseNewhomeProjectUpdateItemInfoResult {
	return poolAlibabaAlihouseNewhomeProjectUpdateItemInfoResult.Get().(*AlibabaAlihouseNewhomeProjectUpdateItemInfoResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectUpdateItemInfoResult 释放AlibabaAlihouseNewhomeProjectUpdateItemInfoResult
func ReleaseAlibabaAlihouseNewhomeProjectUpdateItemInfoResult(v *AlibabaAlihouseNewhomeProjectUpdateItemInfoResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectUpdateItemInfoResult.Put(v)
}
