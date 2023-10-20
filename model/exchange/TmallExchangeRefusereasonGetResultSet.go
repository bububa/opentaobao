package exchange

import (
	"sync"
)

// TmallExchangeRefusereasonGetResultSet 结构体
type TmallExchangeRefusereasonGetResultSet struct {
	// 拒绝原因列表
	Results []Reason `json:"results,omitempty" xml:"results>reason,omitempty"`
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 拒绝原因总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

var poolTmallExchangeRefusereasonGetResultSet = sync.Pool{
	New: func() any {
		return new(TmallExchangeRefusereasonGetResultSet)
	},
}

// GetTmallExchangeRefusereasonGetResultSet() 从对象池中获取TmallExchangeRefusereasonGetResultSet
func GetTmallExchangeRefusereasonGetResultSet() *TmallExchangeRefusereasonGetResultSet {
	return poolTmallExchangeRefusereasonGetResultSet.Get().(*TmallExchangeRefusereasonGetResultSet)
}

// ReleaseTmallExchangeRefusereasonGetResultSet 释放TmallExchangeRefusereasonGetResultSet
func ReleaseTmallExchangeRefusereasonGetResultSet(v *TmallExchangeRefusereasonGetResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	poolTmallExchangeRefusereasonGetResultSet.Put(v)
}
