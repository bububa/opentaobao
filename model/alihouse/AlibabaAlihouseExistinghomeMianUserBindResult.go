package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeMianUserBindResult 结构体
type AlibabaAlihouseExistinghomeMianUserBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeMianUserBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeMianUserBindResult)
	},
}

// GetAlibabaAlihouseExistinghomeMianUserBindResult() 从对象池中获取AlibabaAlihouseExistinghomeMianUserBindResult
func GetAlibabaAlihouseExistinghomeMianUserBindResult() *AlibabaAlihouseExistinghomeMianUserBindResult {
	return poolAlibabaAlihouseExistinghomeMianUserBindResult.Get().(*AlibabaAlihouseExistinghomeMianUserBindResult)
}

// ReleaseAlibabaAlihouseExistinghomeMianUserBindResult 释放AlibabaAlihouseExistinghomeMianUserBindResult
func ReleaseAlibabaAlihouseExistinghomeMianUserBindResult(v *AlibabaAlihouseExistinghomeMianUserBindResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseExistinghomeMianUserBindResult.Put(v)
}
