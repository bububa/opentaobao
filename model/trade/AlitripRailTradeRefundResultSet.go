package trade

import (
	"sync"
)

// AlitripRailTradeRefundResultSet 结构体
type AlitripRailTradeRefundResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口成功/失败
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlitripRailTradeRefundResultSet = sync.Pool{
	New: func() any {
		return new(AlitripRailTradeRefundResultSet)
	},
}

// GetAlitripRailTradeRefundResultSet() 从对象池中获取AlitripRailTradeRefundResultSet
func GetAlitripRailTradeRefundResultSet() *AlitripRailTradeRefundResultSet {
	return poolAlitripRailTradeRefundResultSet.Get().(*AlitripRailTradeRefundResultSet)
}

// ReleaseAlitripRailTradeRefundResultSet 释放AlitripRailTradeRefundResultSet
func ReleaseAlitripRailTradeRefundResultSet(v *AlitripRailTradeRefundResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Module = false
	poolAlitripRailTradeRefundResultSet.Put(v)
}
