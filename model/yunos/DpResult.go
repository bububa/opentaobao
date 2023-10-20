package yunos

import (
	"sync"
)

// DpResult 结构体
type DpResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDpResult = sync.Pool{
	New: func() any {
		return new(DpResult)
	},
}

// GetDpResult() 从对象池中获取DpResult
func GetDpResult() *DpResult {
	return poolDpResult.Get().(*DpResult)
}

// ReleaseDpResult 释放DpResult
func ReleaseDpResult(v *DpResult) {
	v.Message = ""
	v.Code = 0
	v.Success = false
	poolDpResult.Put(v)
}
