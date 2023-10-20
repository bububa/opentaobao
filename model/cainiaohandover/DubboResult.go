package cainiaohandover

import (
	"sync"
)

// DubboResult 结构体
type DubboResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回数据
	Data *AeopActualCarrierResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 返回数据是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDubboResult = sync.Pool{
	New: func() any {
		return new(DubboResult)
	},
}

// GetDubboResult() 从对象池中获取DubboResult
func GetDubboResult() *DubboResult {
	return poolDubboResult.Get().(*DubboResult)
}

// ReleaseDubboResult 释放DubboResult
func ReleaseDubboResult(v *DubboResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolDubboResult.Put(v)
}
