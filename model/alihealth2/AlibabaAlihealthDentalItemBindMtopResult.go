package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalItemBindMtopResult 结构体
type AlibabaAlihealthDentalItemBindMtopResult struct {
	// model
	StoreItemRelVoList []StoreItemRelVo `json:"store_item_rel_vo_list,omitempty" xml:"store_item_rel_vo_list>store_item_rel_vo,omitempty"`
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalItemBindMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemBindMtopResult)
	},
}

// GetAlibabaAlihealthDentalItemBindMtopResult() 从对象池中获取AlibabaAlihealthDentalItemBindMtopResult
func GetAlibabaAlihealthDentalItemBindMtopResult() *AlibabaAlihealthDentalItemBindMtopResult {
	return poolAlibabaAlihealthDentalItemBindMtopResult.Get().(*AlibabaAlihealthDentalItemBindMtopResult)
}

// ReleaseAlibabaAlihealthDentalItemBindMtopResult 释放AlibabaAlihealthDentalItemBindMtopResult
func ReleaseAlibabaAlihealthDentalItemBindMtopResult(v *AlibabaAlihealthDentalItemBindMtopResult) {
	v.StoreItemRelVoList = v.StoreItemRelVoList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthDentalItemBindMtopResult.Put(v)
}
