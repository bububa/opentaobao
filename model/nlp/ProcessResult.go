package nlp

import (
	"sync"
)

// ProcessResult 结构体
type ProcessResult struct {
	// 返回文本处理内容
	TopResult string `json:"top_result,omitempty" xml:"top_result,omitempty"`
	// 返回结果为true则运行成功，为false则运行失败
	TopStatus bool `json:"top_status,omitempty" xml:"top_status,omitempty"`
}

var poolProcessResult = sync.Pool{
	New: func() any {
		return new(ProcessResult)
	},
}

// GetProcessResult() 从对象池中获取ProcessResult
func GetProcessResult() *ProcessResult {
	return poolProcessResult.Get().(*ProcessResult)
}

// ReleaseProcessResult 释放ProcessResult
func ReleaseProcessResult(v *ProcessResult) {
	v.TopResult = ""
	v.TopStatus = false
	poolProcessResult.Put(v)
}
