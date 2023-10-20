package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeVideoChangestatusResult 结构体
type AlibabaAlihouseNewhomeVideoChangestatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 更新结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeVideoChangestatusResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVideoChangestatusResult)
	},
}

// GetAlibabaAlihouseNewhomeVideoChangestatusResult() 从对象池中获取AlibabaAlihouseNewhomeVideoChangestatusResult
func GetAlibabaAlihouseNewhomeVideoChangestatusResult() *AlibabaAlihouseNewhomeVideoChangestatusResult {
	return poolAlibabaAlihouseNewhomeVideoChangestatusResult.Get().(*AlibabaAlihouseNewhomeVideoChangestatusResult)
}

// ReleaseAlibabaAlihouseNewhomeVideoChangestatusResult 释放AlibabaAlihouseNewhomeVideoChangestatusResult
func ReleaseAlibabaAlihouseNewhomeVideoChangestatusResult(v *AlibabaAlihouseNewhomeVideoChangestatusResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseNewhomeVideoChangestatusResult.Put(v)
}
