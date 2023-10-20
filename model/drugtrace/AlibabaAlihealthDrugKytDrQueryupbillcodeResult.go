package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrQueryupbillcodeResult 结构体
type AlibabaAlihealthDrugKytDrQueryupbillcodeResult struct {
	// model
	ModelList []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrQueryupbillcodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrQueryupbillcodeResult)
	},
}

// GetAlibabaAlihealthDrugKytDrQueryupbillcodeResult() 从对象池中获取AlibabaAlihealthDrugKytDrQueryupbillcodeResult
func GetAlibabaAlihealthDrugKytDrQueryupbillcodeResult() *AlibabaAlihealthDrugKytDrQueryupbillcodeResult {
	return poolAlibabaAlihealthDrugKytDrQueryupbillcodeResult.Get().(*AlibabaAlihealthDrugKytDrQueryupbillcodeResult)
}

// ReleaseAlibabaAlihealthDrugKytDrQueryupbillcodeResult 释放AlibabaAlihealthDrugKytDrQueryupbillcodeResult
func ReleaseAlibabaAlihealthDrugKytDrQueryupbillcodeResult(v *AlibabaAlihealthDrugKytDrQueryupbillcodeResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugKytDrQueryupbillcodeResult.Put(v)
}
