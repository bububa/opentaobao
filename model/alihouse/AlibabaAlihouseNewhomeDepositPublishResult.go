package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeDepositPublishResult 结构体
type AlibabaAlihouseNewhomeDepositPublishResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 预存金淘宝商品id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeDepositPublishResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeDepositPublishResult)
	},
}

// GetAlibabaAlihouseNewhomeDepositPublishResult() 从对象池中获取AlibabaAlihouseNewhomeDepositPublishResult
func GetAlibabaAlihouseNewhomeDepositPublishResult() *AlibabaAlihouseNewhomeDepositPublishResult {
	return poolAlibabaAlihouseNewhomeDepositPublishResult.Get().(*AlibabaAlihouseNewhomeDepositPublishResult)
}

// ReleaseAlibabaAlihouseNewhomeDepositPublishResult 释放AlibabaAlihouseNewhomeDepositPublishResult
func ReleaseAlibabaAlihouseNewhomeDepositPublishResult(v *AlibabaAlihouseNewhomeDepositPublishResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeDepositPublishResult.Put(v)
}
