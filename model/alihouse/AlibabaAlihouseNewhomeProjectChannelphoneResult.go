package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectChannelphoneResult 结构体
type AlibabaAlihouseNewhomeProjectChannelphoneResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectChannelphoneResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectChannelphoneResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectChannelphoneResult() 从对象池中获取AlibabaAlihouseNewhomeProjectChannelphoneResult
func GetAlibabaAlihouseNewhomeProjectChannelphoneResult() *AlibabaAlihouseNewhomeProjectChannelphoneResult {
	return poolAlibabaAlihouseNewhomeProjectChannelphoneResult.Get().(*AlibabaAlihouseNewhomeProjectChannelphoneResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectChannelphoneResult 释放AlibabaAlihouseNewhomeProjectChannelphoneResult
func ReleaseAlibabaAlihouseNewhomeProjectChannelphoneResult(v *AlibabaAlihouseNewhomeProjectChannelphoneResult) {
	v.Code = ""
	v.Message = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectChannelphoneResult.Put(v)
}
