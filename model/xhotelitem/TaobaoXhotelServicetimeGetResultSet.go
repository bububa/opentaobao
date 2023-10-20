package xhotelitem

import (
	"sync"
)

// TaobaoXhotelServicetimeGetResultSet 结构体
type TaobaoXhotelServicetimeGetResultSet struct {
	// firstResult
	FirstResults []ServiceTimeDataDo `json:"first_results,omitempty" xml:"first_results>service_time_data_do,omitempty"`
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

var poolTaobaoXhotelServicetimeGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelServicetimeGetResultSet)
	},
}

// GetTaobaoXhotelServicetimeGetResultSet() 从对象池中获取TaobaoXhotelServicetimeGetResultSet
func GetTaobaoXhotelServicetimeGetResultSet() *TaobaoXhotelServicetimeGetResultSet {
	return poolTaobaoXhotelServicetimeGetResultSet.Get().(*TaobaoXhotelServicetimeGetResultSet)
}

// ReleaseTaobaoXhotelServicetimeGetResultSet 释放TaobaoXhotelServicetimeGetResultSet
func ReleaseTaobaoXhotelServicetimeGetResultSet(v *TaobaoXhotelServicetimeGetResultSet) {
	v.FirstResults = v.FirstResults[:0]
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.WarnMessage = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoXhotelServicetimeGetResultSet.Put(v)
}
