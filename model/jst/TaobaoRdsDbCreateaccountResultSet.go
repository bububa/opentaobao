package jst

import (
	"sync"
)

// TaobaoRdsDbCreateaccountResultSet 结构体
type TaobaoRdsDbCreateaccountResultSet struct {
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

var poolTaobaoRdsDbCreateaccountResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoRdsDbCreateaccountResultSet)
	},
}

// GetTaobaoRdsDbCreateaccountResultSet() 从对象池中获取TaobaoRdsDbCreateaccountResultSet
func GetTaobaoRdsDbCreateaccountResultSet() *TaobaoRdsDbCreateaccountResultSet {
	return poolTaobaoRdsDbCreateaccountResultSet.Get().(*TaobaoRdsDbCreateaccountResultSet)
}

// ReleaseTaobaoRdsDbCreateaccountResultSet 释放TaobaoRdsDbCreateaccountResultSet
func ReleaseTaobaoRdsDbCreateaccountResultSet(v *TaobaoRdsDbCreateaccountResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	poolTaobaoRdsDbCreateaccountResultSet.Put(v)
}
