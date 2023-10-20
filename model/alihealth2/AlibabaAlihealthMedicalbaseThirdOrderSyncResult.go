package alihealth2

import (
	"sync"
)

// AlibabaAlihealthMedicalbaseThirdOrderSyncResult 结构体
type AlibabaAlihealthMedicalbaseThirdOrderSyncResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 响应数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 执行是否成功 true成功 false 失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthMedicalbaseThirdOrderSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseThirdOrderSyncResult)
	},
}

// GetAlibabaAlihealthMedicalbaseThirdOrderSyncResult() 从对象池中获取AlibabaAlihealthMedicalbaseThirdOrderSyncResult
func GetAlibabaAlihealthMedicalbaseThirdOrderSyncResult() *AlibabaAlihealthMedicalbaseThirdOrderSyncResult {
	return poolAlibabaAlihealthMedicalbaseThirdOrderSyncResult.Get().(*AlibabaAlihealthMedicalbaseThirdOrderSyncResult)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncResult 释放AlibabaAlihealthMedicalbaseThirdOrderSyncResult
func ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncResult(v *AlibabaAlihealthMedicalbaseThirdOrderSyncResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihealthMedicalbaseThirdOrderSyncResult.Put(v)
}
