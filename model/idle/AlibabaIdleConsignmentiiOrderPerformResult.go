package idle

import (
	"sync"
)

// AlibabaIdleConsignmentiiOrderPerformResult 结构体
type AlibabaIdleConsignmentiiOrderPerformResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleConsignmentiiOrderPerformResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentiiOrderPerformResult)
	},
}

// GetAlibabaIdleConsignmentiiOrderPerformResult() 从对象池中获取AlibabaIdleConsignmentiiOrderPerformResult
func GetAlibabaIdleConsignmentiiOrderPerformResult() *AlibabaIdleConsignmentiiOrderPerformResult {
	return poolAlibabaIdleConsignmentiiOrderPerformResult.Get().(*AlibabaIdleConsignmentiiOrderPerformResult)
}

// ReleaseAlibabaIdleConsignmentiiOrderPerformResult 释放AlibabaIdleConsignmentiiOrderPerformResult
func ReleaseAlibabaIdleConsignmentiiOrderPerformResult(v *AlibabaIdleConsignmentiiOrderPerformResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaIdleConsignmentiiOrderPerformResult.Put(v)
}
