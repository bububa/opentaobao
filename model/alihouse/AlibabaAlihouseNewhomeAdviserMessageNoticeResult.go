package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeAdviserMessageNoticeResult 结构体
type AlibabaAlihouseNewhomeAdviserMessageNoticeResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或失败
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeAdviserMessageNoticeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAdviserMessageNoticeResult)
	},
}

// GetAlibabaAlihouseNewhomeAdviserMessageNoticeResult() 从对象池中获取AlibabaAlihouseNewhomeAdviserMessageNoticeResult
func GetAlibabaAlihouseNewhomeAdviserMessageNoticeResult() *AlibabaAlihouseNewhomeAdviserMessageNoticeResult {
	return poolAlibabaAlihouseNewhomeAdviserMessageNoticeResult.Get().(*AlibabaAlihouseNewhomeAdviserMessageNoticeResult)
}

// ReleaseAlibabaAlihouseNewhomeAdviserMessageNoticeResult 释放AlibabaAlihouseNewhomeAdviserMessageNoticeResult
func ReleaseAlibabaAlihouseNewhomeAdviserMessageNoticeResult(v *AlibabaAlihouseNewhomeAdviserMessageNoticeResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeAdviserMessageNoticeResult.Put(v)
}
