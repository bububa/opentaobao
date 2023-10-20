package nlp

import (
	"sync"
)

// WordResult 结构体
type WordResult struct {
	// 返回文本处理内容
	TopResult string `json:"top_result,omitempty" xml:"top_result,omitempty"`
	// 返回结果为true则运行成功，为false则运行失败
	TopStatus bool `json:"top_status,omitempty" xml:"top_status,omitempty"`
}

var poolWordResult = sync.Pool{
	New: func() any {
		return new(WordResult)
	},
}

// GetWordResult() 从对象池中获取WordResult
func GetWordResult() *WordResult {
	return poolWordResult.Get().(*WordResult)
}

// ReleaseWordResult 释放WordResult
func ReleaseWordResult(v *WordResult) {
	v.TopResult = ""
	v.TopStatus = false
	poolWordResult.Put(v)
}
