package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeSubAccountBindResult 结构体
type AlibabaAlihouseExistinghomeSubAccountBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 入驻是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeSubAccountBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeSubAccountBindResult)
	},
}

// GetAlibabaAlihouseExistinghomeSubAccountBindResult() 从对象池中获取AlibabaAlihouseExistinghomeSubAccountBindResult
func GetAlibabaAlihouseExistinghomeSubAccountBindResult() *AlibabaAlihouseExistinghomeSubAccountBindResult {
	return poolAlibabaAlihouseExistinghomeSubAccountBindResult.Get().(*AlibabaAlihouseExistinghomeSubAccountBindResult)
}

// ReleaseAlibabaAlihouseExistinghomeSubAccountBindResult 释放AlibabaAlihouseExistinghomeSubAccountBindResult
func ReleaseAlibabaAlihouseExistinghomeSubAccountBindResult(v *AlibabaAlihouseExistinghomeSubAccountBindResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseExistinghomeSubAccountBindResult.Put(v)
}
