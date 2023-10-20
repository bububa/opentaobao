package tmallcar

import (
	"sync"
)

// AliAutoResult 结构体
type AliAutoResult struct {
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果
	Data *DaSouEticketVerifyResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliAutoResult = sync.Pool{
	New: func() any {
		return new(AliAutoResult)
	},
}

// GetAliAutoResult() 从对象池中获取AliAutoResult
func GetAliAutoResult() *AliAutoResult {
	return poolAliAutoResult.Get().(*AliAutoResult)
}

// ReleaseAliAutoResult 释放AliAutoResult
func ReleaseAliAutoResult(v *AliAutoResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Data = nil
	v.Success = false
	poolAliAutoResult.Put(v)
}
