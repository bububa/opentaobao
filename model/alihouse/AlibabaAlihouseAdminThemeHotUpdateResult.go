package alihouse

import (
	"sync"
)

// AlibabaAlihouseAdminThemeHotUpdateResult 结构体
type AlibabaAlihouseAdminThemeHotUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseAdminThemeHotUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeHotUpdateResult)
	},
}

// GetAlibabaAlihouseAdminThemeHotUpdateResult() 从对象池中获取AlibabaAlihouseAdminThemeHotUpdateResult
func GetAlibabaAlihouseAdminThemeHotUpdateResult() *AlibabaAlihouseAdminThemeHotUpdateResult {
	return poolAlibabaAlihouseAdminThemeHotUpdateResult.Get().(*AlibabaAlihouseAdminThemeHotUpdateResult)
}

// ReleaseAlibabaAlihouseAdminThemeHotUpdateResult 释放AlibabaAlihouseAdminThemeHotUpdateResult
func ReleaseAlibabaAlihouseAdminThemeHotUpdateResult(v *AlibabaAlihouseAdminThemeHotUpdateResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseAdminThemeHotUpdateResult.Put(v)
}
