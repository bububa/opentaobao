package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQueryactivetimeResult 结构体
type AlibabaAlihealthDrugKytQueryactivetimeResult struct {
	// 码激活状态DTO
	Models []CodeActiveStatusDto `json:"models,omitempty" xml:"models>code_active_status_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQueryactivetimeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryactivetimeResult)
	},
}

// GetAlibabaAlihealthDrugKytQueryactivetimeResult() 从对象池中获取AlibabaAlihealthDrugKytQueryactivetimeResult
func GetAlibabaAlihealthDrugKytQueryactivetimeResult() *AlibabaAlihealthDrugKytQueryactivetimeResult {
	return poolAlibabaAlihealthDrugKytQueryactivetimeResult.Get().(*AlibabaAlihealthDrugKytQueryactivetimeResult)
}

// ReleaseAlibabaAlihealthDrugKytQueryactivetimeResult 释放AlibabaAlihealthDrugKytQueryactivetimeResult
func ReleaseAlibabaAlihealthDrugKytQueryactivetimeResult(v *AlibabaAlihealthDrugKytQueryactivetimeResult) {
	v.Models = v.Models[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugKytQueryactivetimeResult.Put(v)
}
