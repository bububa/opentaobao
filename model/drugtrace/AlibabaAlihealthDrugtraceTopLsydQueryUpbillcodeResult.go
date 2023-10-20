package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult struct {
	// model
	ModelList []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult
func GetAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult() *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult 释放AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult(v *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult.Put(v)
}
