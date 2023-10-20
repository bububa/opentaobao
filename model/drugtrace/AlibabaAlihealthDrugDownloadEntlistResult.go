package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugDownloadEntlistResult 结构体
type AlibabaAlihealthDrugDownloadEntlistResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *DataEntTaskDto `json:"model,omitempty" xml:"model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugDownloadEntlistResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadEntlistResult)
	},
}

// GetAlibabaAlihealthDrugDownloadEntlistResult() 从对象池中获取AlibabaAlihealthDrugDownloadEntlistResult
func GetAlibabaAlihealthDrugDownloadEntlistResult() *AlibabaAlihealthDrugDownloadEntlistResult {
	return poolAlibabaAlihealthDrugDownloadEntlistResult.Get().(*AlibabaAlihealthDrugDownloadEntlistResult)
}

// ReleaseAlibabaAlihealthDrugDownloadEntlistResult 释放AlibabaAlihealthDrugDownloadEntlistResult
func ReleaseAlibabaAlihealthDrugDownloadEntlistResult(v *AlibabaAlihealthDrugDownloadEntlistResult) {
	v.MsgInfo = ""
	v.Model = nil
	v.HttpStatusCode = 0
	v.Success = false
	poolAlibabaAlihealthDrugDownloadEntlistResult.Put(v)
}
