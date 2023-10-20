package alihealth2

import (
	"sync"
)

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult 结构体
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 响应数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 执行是否成功 true成功 false 失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult)
	},
}

// GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult() 从对象池中获取AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult
func GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult() *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult {
	return poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult.Get().(*AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult 释放AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult
func ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult(v *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncResult.Put(v)
}
