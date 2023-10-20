package alihouse

import (
	"sync"
)

// AlibabaAlihouseAdminThemeUpdateStatusResult 结构体
type AlibabaAlihouseAdminThemeUpdateStatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseAdminThemeUpdateStatusResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeUpdateStatusResult)
	},
}

// GetAlibabaAlihouseAdminThemeUpdateStatusResult() 从对象池中获取AlibabaAlihouseAdminThemeUpdateStatusResult
func GetAlibabaAlihouseAdminThemeUpdateStatusResult() *AlibabaAlihouseAdminThemeUpdateStatusResult {
	return poolAlibabaAlihouseAdminThemeUpdateStatusResult.Get().(*AlibabaAlihouseAdminThemeUpdateStatusResult)
}

// ReleaseAlibabaAlihouseAdminThemeUpdateStatusResult 释放AlibabaAlihouseAdminThemeUpdateStatusResult
func ReleaseAlibabaAlihouseAdminThemeUpdateStatusResult(v *AlibabaAlihouseAdminThemeUpdateStatusResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseAdminThemeUpdateStatusResult.Put(v)
}
