package alihouse

import (
	"sync"
)

// AlibabaAlihouseBusinessActivityQueryResult 结构体
type AlibabaAlihouseBusinessActivityQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回值
	Data *BusinessActivityPackageDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseBusinessActivityQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseBusinessActivityQueryResult)
	},
}

// GetAlibabaAlihouseBusinessActivityQueryResult() 从对象池中获取AlibabaAlihouseBusinessActivityQueryResult
func GetAlibabaAlihouseBusinessActivityQueryResult() *AlibabaAlihouseBusinessActivityQueryResult {
	return poolAlibabaAlihouseBusinessActivityQueryResult.Get().(*AlibabaAlihouseBusinessActivityQueryResult)
}

// ReleaseAlibabaAlihouseBusinessActivityQueryResult 释放AlibabaAlihouseBusinessActivityQueryResult
func ReleaseAlibabaAlihouseBusinessActivityQueryResult(v *AlibabaAlihouseBusinessActivityQueryResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseBusinessActivityQueryResult.Put(v)
}
