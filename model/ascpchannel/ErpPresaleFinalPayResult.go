package ascpchannel

import (
	"sync"
)

// ErpPresaleFinalPayResult 结构体
type ErpPresaleFinalPayResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 服务调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolErpPresaleFinalPayResult = sync.Pool{
	New: func() any {
		return new(ErpPresaleFinalPayResult)
	},
}

// GetErpPresaleFinalPayResult() 从对象池中获取ErpPresaleFinalPayResult
func GetErpPresaleFinalPayResult() *ErpPresaleFinalPayResult {
	return poolErpPresaleFinalPayResult.Get().(*ErpPresaleFinalPayResult)
}

// ReleaseErpPresaleFinalPayResult 释放ErpPresaleFinalPayResult
func ReleaseErpPresaleFinalPayResult(v *ErpPresaleFinalPayResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Success = false
	poolErpPresaleFinalPayResult.Put(v)
}
