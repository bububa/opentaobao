package vaccin

import (
	"sync"
)

// AlibabaAlihealthVaccineRegisterSubmitMtopResult 结构体
type AlibabaAlihealthVaccineRegisterSubmitMtopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误提示
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 是否回传成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthVaccineRegisterSubmitMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthVaccineRegisterSubmitMtopResult)
	},
}

// GetAlibabaAlihealthVaccineRegisterSubmitMtopResult() 从对象池中获取AlibabaAlihealthVaccineRegisterSubmitMtopResult
func GetAlibabaAlihealthVaccineRegisterSubmitMtopResult() *AlibabaAlihealthVaccineRegisterSubmitMtopResult {
	return poolAlibabaAlihealthVaccineRegisterSubmitMtopResult.Get().(*AlibabaAlihealthVaccineRegisterSubmitMtopResult)
}

// ReleaseAlibabaAlihealthVaccineRegisterSubmitMtopResult 释放AlibabaAlihealthVaccineRegisterSubmitMtopResult
func ReleaseAlibabaAlihealthVaccineRegisterSubmitMtopResult(v *AlibabaAlihealthVaccineRegisterSubmitMtopResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthVaccineRegisterSubmitMtopResult.Put(v)
}
