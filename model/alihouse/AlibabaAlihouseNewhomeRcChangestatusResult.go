package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeRcChangestatusResult 结构体
type AlibabaAlihouseNewhomeRcChangestatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeRcChangestatusResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRcChangestatusResult)
	},
}

// GetAlibabaAlihouseNewhomeRcChangestatusResult() 从对象池中获取AlibabaAlihouseNewhomeRcChangestatusResult
func GetAlibabaAlihouseNewhomeRcChangestatusResult() *AlibabaAlihouseNewhomeRcChangestatusResult {
	return poolAlibabaAlihouseNewhomeRcChangestatusResult.Get().(*AlibabaAlihouseNewhomeRcChangestatusResult)
}

// ReleaseAlibabaAlihouseNewhomeRcChangestatusResult 释放AlibabaAlihouseNewhomeRcChangestatusResult
func ReleaseAlibabaAlihouseNewhomeRcChangestatusResult(v *AlibabaAlihouseNewhomeRcChangestatusResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseNewhomeRcChangestatusResult.Put(v)
}
