package security

import (
	"sync"
)

// DiAfiResult 结构体
type DiAfiResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 用户参数
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolDiAfiResult = sync.Pool{
	New: func() any {
		return new(DiAfiResult)
	},
}

// GetDiAfiResult() 从对象池中获取DiAfiResult
func GetDiAfiResult() *DiAfiResult {
	return poolDiAfiResult.Get().(*DiAfiResult)
}

// ReleaseDiAfiResult 释放DiAfiResult
func ReleaseDiAfiResult(v *DiAfiResult) {
	v.Code = ""
	v.Data = ""
	v.Msg = ""
	poolDiAfiResult.Put(v)
}
