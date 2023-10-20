package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccountResult 结构体
type AlibabaAlihouseNewhomeAdviserSubmitAccountResult struct {
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeAdviserSubmitAccountResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAdviserSubmitAccountResult)
	},
}

// GetAlibabaAlihouseNewhomeAdviserSubmitAccountResult() 从对象池中获取AlibabaAlihouseNewhomeAdviserSubmitAccountResult
func GetAlibabaAlihouseNewhomeAdviserSubmitAccountResult() *AlibabaAlihouseNewhomeAdviserSubmitAccountResult {
	return poolAlibabaAlihouseNewhomeAdviserSubmitAccountResult.Get().(*AlibabaAlihouseNewhomeAdviserSubmitAccountResult)
}

// ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountResult 释放AlibabaAlihouseNewhomeAdviserSubmitAccountResult
func ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountResult(v *AlibabaAlihouseNewhomeAdviserSubmitAccountResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeAdviserSubmitAccountResult.Put(v)
}
