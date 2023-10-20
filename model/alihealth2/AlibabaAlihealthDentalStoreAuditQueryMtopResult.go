package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalStoreAuditQueryMtopResult 结构体
type AlibabaAlihealthDentalStoreAuditQueryMtopResult struct {
	// model
	StoreAuditVoList []StoreAuditVo `json:"store_audit_vo_list,omitempty" xml:"store_audit_vo_list>store_audit_vo,omitempty"`
	// msg_code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msg_info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalStoreAuditQueryMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStoreAuditQueryMtopResult)
	},
}

// GetAlibabaAlihealthDentalStoreAuditQueryMtopResult() 从对象池中获取AlibabaAlihealthDentalStoreAuditQueryMtopResult
func GetAlibabaAlihealthDentalStoreAuditQueryMtopResult() *AlibabaAlihealthDentalStoreAuditQueryMtopResult {
	return poolAlibabaAlihealthDentalStoreAuditQueryMtopResult.Get().(*AlibabaAlihealthDentalStoreAuditQueryMtopResult)
}

// ReleaseAlibabaAlihealthDentalStoreAuditQueryMtopResult 释放AlibabaAlihealthDentalStoreAuditQueryMtopResult
func ReleaseAlibabaAlihealthDentalStoreAuditQueryMtopResult(v *AlibabaAlihealthDentalStoreAuditQueryMtopResult) {
	v.StoreAuditVoList = v.StoreAuditVoList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthDentalStoreAuditQueryMtopResult.Put(v)
}
