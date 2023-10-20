package alihouse

import (
	"sync"
)

// AlibabaAlihouseBusinessActivityDeleteResult 结构体
type AlibabaAlihouseBusinessActivityDeleteResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 主体公司ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseBusinessActivityDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseBusinessActivityDeleteResult)
	},
}

// GetAlibabaAlihouseBusinessActivityDeleteResult() 从对象池中获取AlibabaAlihouseBusinessActivityDeleteResult
func GetAlibabaAlihouseBusinessActivityDeleteResult() *AlibabaAlihouseBusinessActivityDeleteResult {
	return poolAlibabaAlihouseBusinessActivityDeleteResult.Get().(*AlibabaAlihouseBusinessActivityDeleteResult)
}

// ReleaseAlibabaAlihouseBusinessActivityDeleteResult 释放AlibabaAlihouseBusinessActivityDeleteResult
func ReleaseAlibabaAlihouseBusinessActivityDeleteResult(v *AlibabaAlihouseBusinessActivityDeleteResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseBusinessActivityDeleteResult.Put(v)
}
