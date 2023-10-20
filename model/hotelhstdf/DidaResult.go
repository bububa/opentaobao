package hotelhstdf

import (
	"sync"
)

// DidaResult 结构体
type DidaResult struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 服务是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDidaResult = sync.Pool{
	New: func() any {
		return new(DidaResult)
	},
}

// GetDidaResult() 从对象池中获取DidaResult
func GetDidaResult() *DidaResult {
	return poolDidaResult.Get().(*DidaResult)
}

// ReleaseDidaResult 释放DidaResult
func ReleaseDidaResult(v *DidaResult) {
	v.ResultCode = ""
	v.ResultMsg = ""
	v.Success = false
	poolDidaResult.Put(v)
}
