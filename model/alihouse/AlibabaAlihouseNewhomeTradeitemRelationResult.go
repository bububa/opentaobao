package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeTradeitemRelationResult 结构体
type AlibabaAlihouseNewhomeTradeitemRelationResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeTradeitemRelationResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeitemRelationResult)
	},
}

// GetAlibabaAlihouseNewhomeTradeitemRelationResult() 从对象池中获取AlibabaAlihouseNewhomeTradeitemRelationResult
func GetAlibabaAlihouseNewhomeTradeitemRelationResult() *AlibabaAlihouseNewhomeTradeitemRelationResult {
	return poolAlibabaAlihouseNewhomeTradeitemRelationResult.Get().(*AlibabaAlihouseNewhomeTradeitemRelationResult)
}

// ReleaseAlibabaAlihouseNewhomeTradeitemRelationResult 释放AlibabaAlihouseNewhomeTradeitemRelationResult
func ReleaseAlibabaAlihouseNewhomeTradeitemRelationResult(v *AlibabaAlihouseNewhomeTradeitemRelationResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeTradeitemRelationResult.Put(v)
}
