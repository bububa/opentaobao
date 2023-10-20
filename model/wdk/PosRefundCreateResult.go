package wdk

import (
	"sync"
)

// PosRefundCreateResult 结构体
type PosRefundCreateResult struct {
	// returnCode
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnMsg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPosRefundCreateResult = sync.Pool{
	New: func() any {
		return new(PosRefundCreateResult)
	},
}

// GetPosRefundCreateResult() 从对象池中获取PosRefundCreateResult
func GetPosRefundCreateResult() *PosRefundCreateResult {
	return poolPosRefundCreateResult.Get().(*PosRefundCreateResult)
}

// ReleasePosRefundCreateResult 释放PosRefundCreateResult
func ReleasePosRefundCreateResult(v *PosRefundCreateResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolPosRefundCreateResult.Put(v)
}
