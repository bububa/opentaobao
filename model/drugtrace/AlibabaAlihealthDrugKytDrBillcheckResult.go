package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrBillcheckResult 结构体
type AlibabaAlihealthDrugKytDrBillcheckResult struct {
	// 服务返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 服务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *AlibabaAlihealthDrugKytDrBillcheckModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrBillcheckResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrBillcheckResult)
	},
}

// GetAlibabaAlihealthDrugKytDrBillcheckResult() 从对象池中获取AlibabaAlihealthDrugKytDrBillcheckResult
func GetAlibabaAlihealthDrugKytDrBillcheckResult() *AlibabaAlihealthDrugKytDrBillcheckResult {
	return poolAlibabaAlihealthDrugKytDrBillcheckResult.Get().(*AlibabaAlihealthDrugKytDrBillcheckResult)
}

// ReleaseAlibabaAlihealthDrugKytDrBillcheckResult 释放AlibabaAlihealthDrugKytDrBillcheckResult
func ReleaseAlibabaAlihealthDrugKytDrBillcheckResult(v *AlibabaAlihealthDrugKytDrBillcheckResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.IsSuccess = false
	poolAlibabaAlihealthDrugKytDrBillcheckResult.Put(v)
}
