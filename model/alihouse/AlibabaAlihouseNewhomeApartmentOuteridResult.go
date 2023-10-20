package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeApartmentOuteridResult 结构体
type AlibabaAlihouseNewhomeApartmentOuteridResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeApartmentOuteridResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeApartmentOuteridResult)
	},
}

// GetAlibabaAlihouseNewhomeApartmentOuteridResult() 从对象池中获取AlibabaAlihouseNewhomeApartmentOuteridResult
func GetAlibabaAlihouseNewhomeApartmentOuteridResult() *AlibabaAlihouseNewhomeApartmentOuteridResult {
	return poolAlibabaAlihouseNewhomeApartmentOuteridResult.Get().(*AlibabaAlihouseNewhomeApartmentOuteridResult)
}

// ReleaseAlibabaAlihouseNewhomeApartmentOuteridResult 释放AlibabaAlihouseNewhomeApartmentOuteridResult
func ReleaseAlibabaAlihouseNewhomeApartmentOuteridResult(v *AlibabaAlihouseNewhomeApartmentOuteridResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeApartmentOuteridResult.Put(v)
}
