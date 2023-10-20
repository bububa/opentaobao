package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeEcodeUpdateResult 结构体
type AlibabaAlihouseNewhomeEcodeUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeEcodeUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeEcodeUpdateResult)
	},
}

// GetAlibabaAlihouseNewhomeEcodeUpdateResult() 从对象池中获取AlibabaAlihouseNewhomeEcodeUpdateResult
func GetAlibabaAlihouseNewhomeEcodeUpdateResult() *AlibabaAlihouseNewhomeEcodeUpdateResult {
	return poolAlibabaAlihouseNewhomeEcodeUpdateResult.Get().(*AlibabaAlihouseNewhomeEcodeUpdateResult)
}

// ReleaseAlibabaAlihouseNewhomeEcodeUpdateResult 释放AlibabaAlihouseNewhomeEcodeUpdateResult
func ReleaseAlibabaAlihouseNewhomeEcodeUpdateResult(v *AlibabaAlihouseNewhomeEcodeUpdateResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseNewhomeEcodeUpdateResult.Put(v)
}
