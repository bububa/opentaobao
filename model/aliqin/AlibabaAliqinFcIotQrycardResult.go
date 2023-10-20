package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotQrycardResult 结构体
type AlibabaAliqinFcIotQrycardResult struct {
	// model
	Models []AlibabaAliqinFcIotQrycardModel `json:"models,omitempty" xml:"models>alibaba_aliqin_fc_iot_qrycard_model,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// true返回成功，false返回失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotQrycardResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotQrycardResult)
	},
}

// GetAlibabaAliqinFcIotQrycardResult() 从对象池中获取AlibabaAliqinFcIotQrycardResult
func GetAlibabaAliqinFcIotQrycardResult() *AlibabaAliqinFcIotQrycardResult {
	return poolAlibabaAliqinFcIotQrycardResult.Get().(*AlibabaAliqinFcIotQrycardResult)
}

// ReleaseAlibabaAliqinFcIotQrycardResult 释放AlibabaAliqinFcIotQrycardResult
func ReleaseAlibabaAliqinFcIotQrycardResult(v *AlibabaAliqinFcIotQrycardResult) {
	v.Models = v.Models[:0]
	v.Code = ""
	v.Success = false
	poolAlibabaAliqinFcIotQrycardResult.Put(v)
}
