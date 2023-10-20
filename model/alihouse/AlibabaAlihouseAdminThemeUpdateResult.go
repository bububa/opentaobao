package alihouse

import (
	"sync"
)

// AlibabaAlihouseAdminThemeUpdateResult 结构体
type AlibabaAlihouseAdminThemeUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseAdminThemeUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeUpdateResult)
	},
}

// GetAlibabaAlihouseAdminThemeUpdateResult() 从对象池中获取AlibabaAlihouseAdminThemeUpdateResult
func GetAlibabaAlihouseAdminThemeUpdateResult() *AlibabaAlihouseAdminThemeUpdateResult {
	return poolAlibabaAlihouseAdminThemeUpdateResult.Get().(*AlibabaAlihouseAdminThemeUpdateResult)
}

// ReleaseAlibabaAlihouseAdminThemeUpdateResult 释放AlibabaAlihouseAdminThemeUpdateResult
func ReleaseAlibabaAlihouseAdminThemeUpdateResult(v *AlibabaAlihouseAdminThemeUpdateResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseAdminThemeUpdateResult.Put(v)
}
