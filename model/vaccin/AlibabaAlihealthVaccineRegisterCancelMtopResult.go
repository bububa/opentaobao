package vaccin

import (
	"sync"
)

// AlibabaAlihealthVaccineRegisterCancelMtopResult 结构体
type AlibabaAlihealthVaccineRegisterCancelMtopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误提示
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 是否取消成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthVaccineRegisterCancelMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthVaccineRegisterCancelMtopResult)
	},
}

// GetAlibabaAlihealthVaccineRegisterCancelMtopResult() 从对象池中获取AlibabaAlihealthVaccineRegisterCancelMtopResult
func GetAlibabaAlihealthVaccineRegisterCancelMtopResult() *AlibabaAlihealthVaccineRegisterCancelMtopResult {
	return poolAlibabaAlihealthVaccineRegisterCancelMtopResult.Get().(*AlibabaAlihealthVaccineRegisterCancelMtopResult)
}

// ReleaseAlibabaAlihealthVaccineRegisterCancelMtopResult 释放AlibabaAlihealthVaccineRegisterCancelMtopResult
func ReleaseAlibabaAlihealthVaccineRegisterCancelMtopResult(v *AlibabaAlihealthVaccineRegisterCancelMtopResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthVaccineRegisterCancelMtopResult.Put(v)
}
