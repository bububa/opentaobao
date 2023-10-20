package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrTransportuploadResult 结构体
type AlibabaAlihealthDrugKytDrTransportuploadResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrTransportuploadResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrTransportuploadResult)
	},
}

// GetAlibabaAlihealthDrugKytDrTransportuploadResult() 从对象池中获取AlibabaAlihealthDrugKytDrTransportuploadResult
func GetAlibabaAlihealthDrugKytDrTransportuploadResult() *AlibabaAlihealthDrugKytDrTransportuploadResult {
	return poolAlibabaAlihealthDrugKytDrTransportuploadResult.Get().(*AlibabaAlihealthDrugKytDrTransportuploadResult)
}

// ReleaseAlibabaAlihealthDrugKytDrTransportuploadResult 释放AlibabaAlihealthDrugKytDrTransportuploadResult
func ReleaseAlibabaAlihealthDrugKytDrTransportuploadResult(v *AlibabaAlihealthDrugKytDrTransportuploadResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDrugKytDrTransportuploadResult.Put(v)
}
