package xhotel

import (
	"sync"
)

// TaobaoXhotelBnbpromoUpdateResultSet 结构体
type TaobaoXhotelBnbpromoUpdateResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelBnbpromoUpdateResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoUpdateResultSet)
	},
}

// GetTaobaoXhotelBnbpromoUpdateResultSet() 从对象池中获取TaobaoXhotelBnbpromoUpdateResultSet
func GetTaobaoXhotelBnbpromoUpdateResultSet() *TaobaoXhotelBnbpromoUpdateResultSet {
	return poolTaobaoXhotelBnbpromoUpdateResultSet.Get().(*TaobaoXhotelBnbpromoUpdateResultSet)
}

// ReleaseTaobaoXhotelBnbpromoUpdateResultSet 释放TaobaoXhotelBnbpromoUpdateResultSet
func ReleaseTaobaoXhotelBnbpromoUpdateResultSet(v *TaobaoXhotelBnbpromoUpdateResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoXhotelBnbpromoUpdateResultSet.Put(v)
}
