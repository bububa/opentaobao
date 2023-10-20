package security

import (
	"sync"
)

// JaqSecondCheckResult 结构体
type JaqSecondCheckResult struct {
	// 验证检查结果 1-验证成功 2-验证失败
	SecondCheckResult int64 `json:"second_check_result,omitempty" xml:"second_check_result,omitempty"`
}

var poolJaqSecondCheckResult = sync.Pool{
	New: func() any {
		return new(JaqSecondCheckResult)
	},
}

// GetJaqSecondCheckResult() 从对象池中获取JaqSecondCheckResult
func GetJaqSecondCheckResult() *JaqSecondCheckResult {
	return poolJaqSecondCheckResult.Get().(*JaqSecondCheckResult)
}

// ReleaseJaqSecondCheckResult 释放JaqSecondCheckResult
func ReleaseJaqSecondCheckResult(v *JaqSecondCheckResult) {
	v.SecondCheckResult = 0
	poolJaqSecondCheckResult.Put(v)
}
