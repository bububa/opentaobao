package jst

import (
	"sync"
)

// TaobaoRdsDbGetdbResultSet 结构体
type TaobaoRdsDbGetdbResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// exception
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// totalResults
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

var poolTaobaoRdsDbGetdbResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoRdsDbGetdbResultSet)
	},
}

// GetTaobaoRdsDbGetdbResultSet() 从对象池中获取TaobaoRdsDbGetdbResultSet
func GetTaobaoRdsDbGetdbResultSet() *TaobaoRdsDbGetdbResultSet {
	return poolTaobaoRdsDbGetdbResultSet.Get().(*TaobaoRdsDbGetdbResultSet)
}

// ReleaseTaobaoRdsDbGetdbResultSet 释放TaobaoRdsDbGetdbResultSet
func ReleaseTaobaoRdsDbGetdbResultSet(v *TaobaoRdsDbGetdbResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	poolTaobaoRdsDbGetdbResultSet.Put(v)
}
