package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeSystemSellerResult 结构体
type AlibabaAlihouseNewhomeSystemSellerResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeSystemSellerResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeSystemSellerResult)
	},
}

// GetAlibabaAlihouseNewhomeSystemSellerResult() 从对象池中获取AlibabaAlihouseNewhomeSystemSellerResult
func GetAlibabaAlihouseNewhomeSystemSellerResult() *AlibabaAlihouseNewhomeSystemSellerResult {
	return poolAlibabaAlihouseNewhomeSystemSellerResult.Get().(*AlibabaAlihouseNewhomeSystemSellerResult)
}

// ReleaseAlibabaAlihouseNewhomeSystemSellerResult 释放AlibabaAlihouseNewhomeSystemSellerResult
func ReleaseAlibabaAlihouseNewhomeSystemSellerResult(v *AlibabaAlihouseNewhomeSystemSellerResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeSystemSellerResult.Put(v)
}
