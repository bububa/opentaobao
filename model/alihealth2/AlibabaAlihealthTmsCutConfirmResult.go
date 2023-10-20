package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTmsCutConfirmResult 结构体
type AlibabaAlihealthTmsCutConfirmResult struct {
	// 结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthTmsCutConfirmResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTmsCutConfirmResult)
	},
}

// GetAlibabaAlihealthTmsCutConfirmResult() 从对象池中获取AlibabaAlihealthTmsCutConfirmResult
func GetAlibabaAlihealthTmsCutConfirmResult() *AlibabaAlihealthTmsCutConfirmResult {
	return poolAlibabaAlihealthTmsCutConfirmResult.Get().(*AlibabaAlihealthTmsCutConfirmResult)
}

// ReleaseAlibabaAlihealthTmsCutConfirmResult 释放AlibabaAlihealthTmsCutConfirmResult
func ReleaseAlibabaAlihealthTmsCutConfirmResult(v *AlibabaAlihealthTmsCutConfirmResult) {
	v.Data = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaAlihealthTmsCutConfirmResult.Put(v)
}
