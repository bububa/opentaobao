package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalBindAuditQueryMtopResult 结构体
type AlibabaAlihealthDentalBindAuditQueryMtopResult struct {
	// model
	StoreitemrelvoList []StoreItemRelVo `json:"storeitemrelvo_list,omitempty" xml:"storeitemrelvo_list>store_item_rel_vo,omitempty"`
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalBindAuditQueryMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalBindAuditQueryMtopResult)
	},
}

// GetAlibabaAlihealthDentalBindAuditQueryMtopResult() 从对象池中获取AlibabaAlihealthDentalBindAuditQueryMtopResult
func GetAlibabaAlihealthDentalBindAuditQueryMtopResult() *AlibabaAlihealthDentalBindAuditQueryMtopResult {
	return poolAlibabaAlihealthDentalBindAuditQueryMtopResult.Get().(*AlibabaAlihealthDentalBindAuditQueryMtopResult)
}

// ReleaseAlibabaAlihealthDentalBindAuditQueryMtopResult 释放AlibabaAlihealthDentalBindAuditQueryMtopResult
func ReleaseAlibabaAlihealthDentalBindAuditQueryMtopResult(v *AlibabaAlihealthDentalBindAuditQueryMtopResult) {
	v.StoreitemrelvoList = v.StoreitemrelvoList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthDentalBindAuditQueryMtopResult.Put(v)
}
