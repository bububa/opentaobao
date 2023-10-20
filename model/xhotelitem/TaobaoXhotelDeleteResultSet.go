package xhotelitem

import (
	"sync"
)

// TaobaoXhotelDeleteResultSet 结构体
type TaobaoXhotelDeleteResultSet struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否出错
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoXhotelDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDeleteResultSet)
	},
}

// GetTaobaoXhotelDeleteResultSet() 从对象池中获取TaobaoXhotelDeleteResultSet
func GetTaobaoXhotelDeleteResultSet() *TaobaoXhotelDeleteResultSet {
	return poolTaobaoXhotelDeleteResultSet.Get().(*TaobaoXhotelDeleteResultSet)
}

// ReleaseTaobaoXhotelDeleteResultSet 释放TaobaoXhotelDeleteResultSet
func ReleaseTaobaoXhotelDeleteResultSet(v *TaobaoXhotelDeleteResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Error = false
	poolTaobaoXhotelDeleteResultSet.Put(v)
}
