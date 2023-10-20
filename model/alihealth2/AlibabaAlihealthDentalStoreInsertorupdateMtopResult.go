package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalStoreInsertorupdateMtopResult 结构体
type AlibabaAlihealthDentalStoreInsertorupdateMtopResult struct {
	// msg_code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msg_info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *StoreAuditVo `json:"model,omitempty" xml:"model,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalStoreInsertorupdateMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStoreInsertorupdateMtopResult)
	},
}

// GetAlibabaAlihealthDentalStoreInsertorupdateMtopResult() 从对象池中获取AlibabaAlihealthDentalStoreInsertorupdateMtopResult
func GetAlibabaAlihealthDentalStoreInsertorupdateMtopResult() *AlibabaAlihealthDentalStoreInsertorupdateMtopResult {
	return poolAlibabaAlihealthDentalStoreInsertorupdateMtopResult.Get().(*AlibabaAlihealthDentalStoreInsertorupdateMtopResult)
}

// ReleaseAlibabaAlihealthDentalStoreInsertorupdateMtopResult 释放AlibabaAlihealthDentalStoreInsertorupdateMtopResult
func ReleaseAlibabaAlihealthDentalStoreInsertorupdateMtopResult(v *AlibabaAlihealthDentalStoreInsertorupdateMtopResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDentalStoreInsertorupdateMtopResult.Put(v)
}
