package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeLinkInfoObtainResult 结构体
type AlibabaAlihouseNewhomeLinkInfoObtainResult struct {
	// 返回链接地址
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// message
	Msg bool `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolAlibabaAlihouseNewhomeLinkInfoObtainResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLinkInfoObtainResult)
	},
}

// GetAlibabaAlihouseNewhomeLinkInfoObtainResult() 从对象池中获取AlibabaAlihouseNewhomeLinkInfoObtainResult
func GetAlibabaAlihouseNewhomeLinkInfoObtainResult() *AlibabaAlihouseNewhomeLinkInfoObtainResult {
	return poolAlibabaAlihouseNewhomeLinkInfoObtainResult.Get().(*AlibabaAlihouseNewhomeLinkInfoObtainResult)
}

// ReleaseAlibabaAlihouseNewhomeLinkInfoObtainResult 释放AlibabaAlihouseNewhomeLinkInfoObtainResult
func ReleaseAlibabaAlihouseNewhomeLinkInfoObtainResult(v *AlibabaAlihouseNewhomeLinkInfoObtainResult) {
	v.Data = ""
	v.Code = ""
	v.IsSuccess = false
	v.Msg = false
	poolAlibabaAlihouseNewhomeLinkInfoObtainResult.Put(v)
}
