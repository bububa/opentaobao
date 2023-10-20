package xhotelitem

import (
	"sync"
)

// TaobaoXhotelServicetimeUpdateResultSet 结构体
type TaobaoXhotelServicetimeUpdateResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// exception
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// warnMessage
	WarnMessage string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// totalResults
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// hasNext
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolTaobaoXhotelServicetimeUpdateResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelServicetimeUpdateResultSet)
	},
}

// GetTaobaoXhotelServicetimeUpdateResultSet() 从对象池中获取TaobaoXhotelServicetimeUpdateResultSet
func GetTaobaoXhotelServicetimeUpdateResultSet() *TaobaoXhotelServicetimeUpdateResultSet {
	return poolTaobaoXhotelServicetimeUpdateResultSet.Get().(*TaobaoXhotelServicetimeUpdateResultSet)
}

// ReleaseTaobaoXhotelServicetimeUpdateResultSet 释放TaobaoXhotelServicetimeUpdateResultSet
func ReleaseTaobaoXhotelServicetimeUpdateResultSet(v *TaobaoXhotelServicetimeUpdateResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.WarnMessage = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoXhotelServicetimeUpdateResultSet.Put(v)
}
