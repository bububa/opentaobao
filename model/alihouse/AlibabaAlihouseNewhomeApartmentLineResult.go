package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeApartmentLineResult 结构体
type AlibabaAlihouseNewhomeApartmentLineResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeApartmentLineResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeApartmentLineResult)
	},
}

// GetAlibabaAlihouseNewhomeApartmentLineResult() 从对象池中获取AlibabaAlihouseNewhomeApartmentLineResult
func GetAlibabaAlihouseNewhomeApartmentLineResult() *AlibabaAlihouseNewhomeApartmentLineResult {
	return poolAlibabaAlihouseNewhomeApartmentLineResult.Get().(*AlibabaAlihouseNewhomeApartmentLineResult)
}

// ReleaseAlibabaAlihouseNewhomeApartmentLineResult 释放AlibabaAlihouseNewhomeApartmentLineResult
func ReleaseAlibabaAlihouseNewhomeApartmentLineResult(v *AlibabaAlihouseNewhomeApartmentLineResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeApartmentLineResult.Put(v)
}
