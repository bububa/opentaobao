package charity

import (
	"sync"
)

// CsrResult 结构体
type CsrResult struct {
	// 接口响应消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态码：200表示正常，非200表示异常
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 公益时授权链接结果
	Data *JumpAddressHsfResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolCsrResult = sync.Pool{
	New: func() any {
		return new(CsrResult)
	},
}

// GetCsrResult() 从对象池中获取CsrResult
func GetCsrResult() *CsrResult {
	return poolCsrResult.Get().(*CsrResult)
}

// ReleaseCsrResult 释放CsrResult
func ReleaseCsrResult(v *CsrResult) {
	v.Msg = ""
	v.Code = 0
	v.Data = nil
	poolCsrResult.Put(v)
}
