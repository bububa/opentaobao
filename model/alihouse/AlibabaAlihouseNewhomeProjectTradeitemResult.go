package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectTradeitemResult 结构体
type AlibabaAlihouseNewhomeProjectTradeitemResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectTradeitemResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTradeitemResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectTradeitemResult() 从对象池中获取AlibabaAlihouseNewhomeProjectTradeitemResult
func GetAlibabaAlihouseNewhomeProjectTradeitemResult() *AlibabaAlihouseNewhomeProjectTradeitemResult {
	return poolAlibabaAlihouseNewhomeProjectTradeitemResult.Get().(*AlibabaAlihouseNewhomeProjectTradeitemResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectTradeitemResult 释放AlibabaAlihouseNewhomeProjectTradeitemResult
func ReleaseAlibabaAlihouseNewhomeProjectTradeitemResult(v *AlibabaAlihouseNewhomeProjectTradeitemResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectTradeitemResult.Put(v)
}
