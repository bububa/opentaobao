package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBnbownerDeleteResultSet 结构体
type TaobaoXhotelBnbownerDeleteResultSet struct {
	// 系统自动生成
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否出错
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoXhotelBnbownerDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbownerDeleteResultSet)
	},
}

// GetTaobaoXhotelBnbownerDeleteResultSet() 从对象池中获取TaobaoXhotelBnbownerDeleteResultSet
func GetTaobaoXhotelBnbownerDeleteResultSet() *TaobaoXhotelBnbownerDeleteResultSet {
	return poolTaobaoXhotelBnbownerDeleteResultSet.Get().(*TaobaoXhotelBnbownerDeleteResultSet)
}

// ReleaseTaobaoXhotelBnbownerDeleteResultSet 释放TaobaoXhotelBnbownerDeleteResultSet
func ReleaseTaobaoXhotelBnbownerDeleteResultSet(v *TaobaoXhotelBnbownerDeleteResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Error = false
	poolTaobaoXhotelBnbownerDeleteResultSet.Put(v)
}
