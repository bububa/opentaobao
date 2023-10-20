package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYbGetcoderelationResult 结构体
type AlibabaAlihealthDrugKytYbGetcoderelationResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	ModelList *AlibabaAlihealthDrugKytYbGetcoderelationModel `json:"model_list,omitempty" xml:"model_list,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytYbGetcoderelationResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYbGetcoderelationResult)
	},
}

// GetAlibabaAlihealthDrugKytYbGetcoderelationResult() 从对象池中获取AlibabaAlihealthDrugKytYbGetcoderelationResult
func GetAlibabaAlihealthDrugKytYbGetcoderelationResult() *AlibabaAlihealthDrugKytYbGetcoderelationResult {
	return poolAlibabaAlihealthDrugKytYbGetcoderelationResult.Get().(*AlibabaAlihealthDrugKytYbGetcoderelationResult)
}

// ReleaseAlibabaAlihealthDrugKytYbGetcoderelationResult 释放AlibabaAlihealthDrugKytYbGetcoderelationResult
func ReleaseAlibabaAlihealthDrugKytYbGetcoderelationResult(v *AlibabaAlihealthDrugKytYbGetcoderelationResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ModelList = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytYbGetcoderelationResult.Put(v)
}
