package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult struct {
	// model
	ModelList []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult
func GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult 释放AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult(v *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult.Put(v)
}
