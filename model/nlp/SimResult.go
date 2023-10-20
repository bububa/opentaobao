package nlp

import (
	"sync"
)

// SimResult 结构体
type SimResult struct {
	// 返回文本处理内容
	TopResult string `json:"top_result,omitempty" xml:"top_result,omitempty"`
	// 返回结果为true则运行成功，为false则运行失败
	TopStatus bool `json:"top_status,omitempty" xml:"top_status,omitempty"`
}

var poolSimResult = sync.Pool{
	New: func() any {
		return new(SimResult)
	},
}

// GetSimResult() 从对象池中获取SimResult
func GetSimResult() *SimResult {
	return poolSimResult.Get().(*SimResult)
}

// ReleaseSimResult 释放SimResult
func ReleaseSimResult(v *SimResult) {
	v.TopResult = ""
	v.TopStatus = false
	poolSimResult.Put(v)
}
