package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetproteminfoResult 结构体
type AlibabaAlihealthDrugKytDrGetproteminfoResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *AlibabaAlihealthDrugKytDrGetproteminfoModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetproteminfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetproteminfoResult)
	},
}

// GetAlibabaAlihealthDrugKytDrGetproteminfoResult() 从对象池中获取AlibabaAlihealthDrugKytDrGetproteminfoResult
func GetAlibabaAlihealthDrugKytDrGetproteminfoResult() *AlibabaAlihealthDrugKytDrGetproteminfoResult {
	return poolAlibabaAlihealthDrugKytDrGetproteminfoResult.Get().(*AlibabaAlihealthDrugKytDrGetproteminfoResult)
}

// ReleaseAlibabaAlihealthDrugKytDrGetproteminfoResult 释放AlibabaAlihealthDrugKytDrGetproteminfoResult
func ReleaseAlibabaAlihealthDrugKytDrGetproteminfoResult(v *AlibabaAlihealthDrugKytDrGetproteminfoResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytDrGetproteminfoResult.Put(v)
}
