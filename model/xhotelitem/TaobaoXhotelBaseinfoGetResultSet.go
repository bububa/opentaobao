package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBaseinfoGetResultSet 结构体
type TaobaoXhotelBaseinfoGetResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 酒店基础信息
	XhotelBaseInfo *XHotelBaseInfo `json:"xhotel_base_info,omitempty" xml:"xhotel_base_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelBaseinfoGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBaseinfoGetResultSet)
	},
}

// GetTaobaoXhotelBaseinfoGetResultSet() 从对象池中获取TaobaoXhotelBaseinfoGetResultSet
func GetTaobaoXhotelBaseinfoGetResultSet() *TaobaoXhotelBaseinfoGetResultSet {
	return poolTaobaoXhotelBaseinfoGetResultSet.Get().(*TaobaoXhotelBaseinfoGetResultSet)
}

// ReleaseTaobaoXhotelBaseinfoGetResultSet 释放TaobaoXhotelBaseinfoGetResultSet
func ReleaseTaobaoXhotelBaseinfoGetResultSet(v *TaobaoXhotelBaseinfoGetResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.XhotelBaseInfo = nil
	v.Success = false
	poolTaobaoXhotelBaseinfoGetResultSet.Put(v)
}
