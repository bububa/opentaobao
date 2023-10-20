package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalItemUnbindMtopResult 结构体
type AlibabaAlihealthDentalItemUnbindMtopResult struct {
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalItemUnbindMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemUnbindMtopResult)
	},
}

// GetAlibabaAlihealthDentalItemUnbindMtopResult() 从对象池中获取AlibabaAlihealthDentalItemUnbindMtopResult
func GetAlibabaAlihealthDentalItemUnbindMtopResult() *AlibabaAlihealthDentalItemUnbindMtopResult {
	return poolAlibabaAlihealthDentalItemUnbindMtopResult.Get().(*AlibabaAlihealthDentalItemUnbindMtopResult)
}

// ReleaseAlibabaAlihealthDentalItemUnbindMtopResult 释放AlibabaAlihealthDentalItemUnbindMtopResult
func ReleaseAlibabaAlihealthDentalItemUnbindMtopResult(v *AlibabaAlihealthDentalItemUnbindMtopResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDentalItemUnbindMtopResult.Put(v)
}
