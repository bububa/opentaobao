package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrAssociateequiResult 结构体
type AlibabaAlihealthDrugKytDrAssociateequiResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrAssociateequiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrAssociateequiResult)
	},
}

// GetAlibabaAlihealthDrugKytDrAssociateequiResult() 从对象池中获取AlibabaAlihealthDrugKytDrAssociateequiResult
func GetAlibabaAlihealthDrugKytDrAssociateequiResult() *AlibabaAlihealthDrugKytDrAssociateequiResult {
	return poolAlibabaAlihealthDrugKytDrAssociateequiResult.Get().(*AlibabaAlihealthDrugKytDrAssociateequiResult)
}

// ReleaseAlibabaAlihealthDrugKytDrAssociateequiResult 释放AlibabaAlihealthDrugKytDrAssociateequiResult
func ReleaseAlibabaAlihealthDrugKytDrAssociateequiResult(v *AlibabaAlihealthDrugKytDrAssociateequiResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDrugKytDrAssociateequiResult.Put(v)
}
