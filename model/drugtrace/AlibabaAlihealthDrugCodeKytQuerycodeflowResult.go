package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflowResult 结构体
type AlibabaAlihealthDrugCodeKytQuerycodeflowResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *AlibabaAlihealthDrugCodeKytQuerycodeflowModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeflowResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytQuerycodeflowResult)
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeflowResult() 从对象池中获取AlibabaAlihealthDrugCodeKytQuerycodeflowResult
func GetAlibabaAlihealthDrugCodeKytQuerycodeflowResult() *AlibabaAlihealthDrugCodeKytQuerycodeflowResult {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeflowResult.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeflowResult)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowResult 释放AlibabaAlihealthDrugCodeKytQuerycodeflowResult
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowResult(v *AlibabaAlihealthDrugCodeKytQuerycodeflowResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugCodeKytQuerycodeflowResult.Put(v)
}
