package alihouse

import (
	"sync"
)

// AlibabaAlihouseAdminThemeCreateResult 结构体
type AlibabaAlihouseAdminThemeCreateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseAdminThemeCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeCreateResult)
	},
}

// GetAlibabaAlihouseAdminThemeCreateResult() 从对象池中获取AlibabaAlihouseAdminThemeCreateResult
func GetAlibabaAlihouseAdminThemeCreateResult() *AlibabaAlihouseAdminThemeCreateResult {
	return poolAlibabaAlihouseAdminThemeCreateResult.Get().(*AlibabaAlihouseAdminThemeCreateResult)
}

// ReleaseAlibabaAlihouseAdminThemeCreateResult 释放AlibabaAlihouseAdminThemeCreateResult
func ReleaseAlibabaAlihouseAdminThemeCreateResult(v *AlibabaAlihouseAdminThemeCreateResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseAdminThemeCreateResult.Put(v)
}
