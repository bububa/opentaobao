package xhotelonlineorder

import (
	"sync"
)

// TaobaoXhotelPmsGuestbillGetVtwoResultSet 结构体
type TaobaoXhotelPmsGuestbillGetVtwoResultSet struct {
	// 账单列表中涉及到的金额费用单位均为分
	Results []OrderBillInfo `json:"results,omitempty" xml:"results>order_bill_info,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelPmsGuestbillGetVtwoResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelPmsGuestbillGetVtwoResultSet)
	},
}

// GetTaobaoXhotelPmsGuestbillGetVtwoResultSet() 从对象池中获取TaobaoXhotelPmsGuestbillGetVtwoResultSet
func GetTaobaoXhotelPmsGuestbillGetVtwoResultSet() *TaobaoXhotelPmsGuestbillGetVtwoResultSet {
	return poolTaobaoXhotelPmsGuestbillGetVtwoResultSet.Get().(*TaobaoXhotelPmsGuestbillGetVtwoResultSet)
}

// ReleaseTaobaoXhotelPmsGuestbillGetVtwoResultSet 释放TaobaoXhotelPmsGuestbillGetVtwoResultSet
func ReleaseTaobaoXhotelPmsGuestbillGetVtwoResultSet(v *TaobaoXhotelPmsGuestbillGetVtwoResultSet) {
	v.Results = v.Results[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTaobaoXhotelPmsGuestbillGetVtwoResultSet.Put(v)
}
