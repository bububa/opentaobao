package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBnbpromoDeleteResultSet 结构体
type TaobaoXhotelBnbpromoDeleteResultSet struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelBnbpromoDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoDeleteResultSet)
	},
}

// GetTaobaoXhotelBnbpromoDeleteResultSet() 从对象池中获取TaobaoXhotelBnbpromoDeleteResultSet
func GetTaobaoXhotelBnbpromoDeleteResultSet() *TaobaoXhotelBnbpromoDeleteResultSet {
	return poolTaobaoXhotelBnbpromoDeleteResultSet.Get().(*TaobaoXhotelBnbpromoDeleteResultSet)
}

// ReleaseTaobaoXhotelBnbpromoDeleteResultSet 释放TaobaoXhotelBnbpromoDeleteResultSet
func ReleaseTaobaoXhotelBnbpromoDeleteResultSet(v *TaobaoXhotelBnbpromoDeleteResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoXhotelBnbpromoDeleteResultSet.Put(v)
}
