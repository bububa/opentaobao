package exchange

import (
	"sync"
)

// TmallExchangeMessageAddResultSet 结构体
type TmallExchangeMessageAddResultSet struct {
	// 留言信息
	Results []RefundMessage `json:"results,omitempty" xml:"results>refund_message,omitempty"`
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolTmallExchangeMessageAddResultSet = sync.Pool{
	New: func() any {
		return new(TmallExchangeMessageAddResultSet)
	},
}

// GetTmallExchangeMessageAddResultSet() 从对象池中获取TmallExchangeMessageAddResultSet
func GetTmallExchangeMessageAddResultSet() *TmallExchangeMessageAddResultSet {
	return poolTmallExchangeMessageAddResultSet.Get().(*TmallExchangeMessageAddResultSet)
}

// ReleaseTmallExchangeMessageAddResultSet 释放TmallExchangeMessageAddResultSet
func ReleaseTmallExchangeMessageAddResultSet(v *TmallExchangeMessageAddResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	poolTmallExchangeMessageAddResultSet.Put(v)
}
