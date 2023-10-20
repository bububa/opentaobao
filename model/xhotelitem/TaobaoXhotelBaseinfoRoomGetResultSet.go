package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBaseinfoRoomGetResultSet 结构体
type TaobaoXhotelBaseinfoRoomGetResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 酒店基础信息
	XhotelBaseInfo *XHotelInfoWithRoom `json:"xhotel_base_info,omitempty" xml:"xhotel_base_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelBaseinfoRoomGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBaseinfoRoomGetResultSet)
	},
}

// GetTaobaoXhotelBaseinfoRoomGetResultSet() 从对象池中获取TaobaoXhotelBaseinfoRoomGetResultSet
func GetTaobaoXhotelBaseinfoRoomGetResultSet() *TaobaoXhotelBaseinfoRoomGetResultSet {
	return poolTaobaoXhotelBaseinfoRoomGetResultSet.Get().(*TaobaoXhotelBaseinfoRoomGetResultSet)
}

// ReleaseTaobaoXhotelBaseinfoRoomGetResultSet 释放TaobaoXhotelBaseinfoRoomGetResultSet
func ReleaseTaobaoXhotelBaseinfoRoomGetResultSet(v *TaobaoXhotelBaseinfoRoomGetResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.XhotelBaseInfo = nil
	v.Success = false
	poolTaobaoXhotelBaseinfoRoomGetResultSet.Put(v)
}
