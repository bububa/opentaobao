package idle

import (
	"sync"
)

// AlibabaIdleConsignmentOrderPerformResult 结构体
type AlibabaIdleConsignmentOrderPerformResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleConsignmentOrderPerformResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentOrderPerformResult)
	},
}

// GetAlibabaIdleConsignmentOrderPerformResult() 从对象池中获取AlibabaIdleConsignmentOrderPerformResult
func GetAlibabaIdleConsignmentOrderPerformResult() *AlibabaIdleConsignmentOrderPerformResult {
	return poolAlibabaIdleConsignmentOrderPerformResult.Get().(*AlibabaIdleConsignmentOrderPerformResult)
}

// ReleaseAlibabaIdleConsignmentOrderPerformResult 释放AlibabaIdleConsignmentOrderPerformResult
func ReleaseAlibabaIdleConsignmentOrderPerformResult(v *AlibabaIdleConsignmentOrderPerformResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaIdleConsignmentOrderPerformResult.Put(v)
}
