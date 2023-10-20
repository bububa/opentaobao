package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult 结构体
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult() 从对象池中获取AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult
func GetAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult() *AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult {
	return poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult.Get().(*AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult 释放AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult
func ReleaseAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult(v *AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncResult.Put(v)
}
