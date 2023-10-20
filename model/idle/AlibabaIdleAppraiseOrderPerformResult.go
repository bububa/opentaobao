package idle

import (
	"sync"
)

// AlibabaIdleAppraiseOrderPerformResult 结构体
type AlibabaIdleAppraiseOrderPerformResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolAlibabaIdleAppraiseOrderPerformResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAppraiseOrderPerformResult)
	},
}

// GetAlibabaIdleAppraiseOrderPerformResult() 从对象池中获取AlibabaIdleAppraiseOrderPerformResult
func GetAlibabaIdleAppraiseOrderPerformResult() *AlibabaIdleAppraiseOrderPerformResult {
	return poolAlibabaIdleAppraiseOrderPerformResult.Get().(*AlibabaIdleAppraiseOrderPerformResult)
}

// ReleaseAlibabaIdleAppraiseOrderPerformResult 释放AlibabaIdleAppraiseOrderPerformResult
func ReleaseAlibabaIdleAppraiseOrderPerformResult(v *AlibabaIdleAppraiseOrderPerformResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Succ = false
	poolAlibabaIdleAppraiseOrderPerformResult.Put(v)
}
