package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetupteminfoResult 结构体
type AlibabaAlihealthDrugKytDrGetupteminfoResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *AlibabaAlihealthDrugKytDrGetupteminfoModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetupteminfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetupteminfoResult)
	},
}

// GetAlibabaAlihealthDrugKytDrGetupteminfoResult() 从对象池中获取AlibabaAlihealthDrugKytDrGetupteminfoResult
func GetAlibabaAlihealthDrugKytDrGetupteminfoResult() *AlibabaAlihealthDrugKytDrGetupteminfoResult {
	return poolAlibabaAlihealthDrugKytDrGetupteminfoResult.Get().(*AlibabaAlihealthDrugKytDrGetupteminfoResult)
}

// ReleaseAlibabaAlihealthDrugKytDrGetupteminfoResult 释放AlibabaAlihealthDrugKytDrGetupteminfoResult
func ReleaseAlibabaAlihealthDrugKytDrGetupteminfoResult(v *AlibabaAlihealthDrugKytDrGetupteminfoResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytDrGetupteminfoResult.Put(v)
}
