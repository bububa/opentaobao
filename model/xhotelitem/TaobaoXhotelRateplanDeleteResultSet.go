package xhotelitem

import (
	"sync"
)

// TaobaoXhotelRateplanDeleteResultSet 结构体
type TaobaoXhotelRateplanDeleteResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 房价id
	Rpid string `json:"rpid,omitempty" xml:"rpid,omitempty"`
}

var poolTaobaoXhotelRateplanDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateplanDeleteResultSet)
	},
}

// GetTaobaoXhotelRateplanDeleteResultSet() 从对象池中获取TaobaoXhotelRateplanDeleteResultSet
func GetTaobaoXhotelRateplanDeleteResultSet() *TaobaoXhotelRateplanDeleteResultSet {
	return poolTaobaoXhotelRateplanDeleteResultSet.Get().(*TaobaoXhotelRateplanDeleteResultSet)
}

// ReleaseTaobaoXhotelRateplanDeleteResultSet 释放TaobaoXhotelRateplanDeleteResultSet
func ReleaseTaobaoXhotelRateplanDeleteResultSet(v *TaobaoXhotelRateplanDeleteResultSet) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Rpid = ""
	poolTaobaoXhotelRateplanDeleteResultSet.Put(v)
}
