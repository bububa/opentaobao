package xhotelitem

import (
	"sync"
)

// TaobaoXhotelRoomtypeDeletePublicResultSet 结构体
type TaobaoXhotelRoomtypeDeletePublicResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// deleteResult
	DeleteResult string `json:"delete_result,omitempty" xml:"delete_result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelRoomtypeDeletePublicResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomtypeDeletePublicResultSet)
	},
}

// GetTaobaoXhotelRoomtypeDeletePublicResultSet() 从对象池中获取TaobaoXhotelRoomtypeDeletePublicResultSet
func GetTaobaoXhotelRoomtypeDeletePublicResultSet() *TaobaoXhotelRoomtypeDeletePublicResultSet {
	return poolTaobaoXhotelRoomtypeDeletePublicResultSet.Get().(*TaobaoXhotelRoomtypeDeletePublicResultSet)
}

// ReleaseTaobaoXhotelRoomtypeDeletePublicResultSet 释放TaobaoXhotelRoomtypeDeletePublicResultSet
func ReleaseTaobaoXhotelRoomtypeDeletePublicResultSet(v *TaobaoXhotelRoomtypeDeletePublicResultSet) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.DeleteResult = ""
	v.Success = false
	poolTaobaoXhotelRoomtypeDeletePublicResultSet.Put(v)
}
