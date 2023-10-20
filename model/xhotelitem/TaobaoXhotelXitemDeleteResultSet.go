package xhotelitem

import (
	"sync"
)

// TaobaoXhotelXitemDeleteResultSet 结构体
type TaobaoXhotelXitemDeleteResultSet struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

var poolTaobaoXhotelXitemDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelXitemDeleteResultSet)
	},
}

// GetTaobaoXhotelXitemDeleteResultSet() 从对象池中获取TaobaoXhotelXitemDeleteResultSet
func GetTaobaoXhotelXitemDeleteResultSet() *TaobaoXhotelXitemDeleteResultSet {
	return poolTaobaoXhotelXitemDeleteResultSet.Get().(*TaobaoXhotelXitemDeleteResultSet)
}

// ReleaseTaobaoXhotelXitemDeleteResultSet 释放TaobaoXhotelXitemDeleteResultSet
func ReleaseTaobaoXhotelXitemDeleteResultSet(v *TaobaoXhotelXitemDeleteResultSet) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	poolTaobaoXhotelXitemDeleteResultSet.Put(v)
}
