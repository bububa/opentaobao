package rail

import (
	"sync"
)

// AlitripRailTradeCloseticketResultSet 结构体
type AlitripRailTradeCloseticketResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功失败
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlitripRailTradeCloseticketResultSet = sync.Pool{
	New: func() any {
		return new(AlitripRailTradeCloseticketResultSet)
	},
}

// GetAlitripRailTradeCloseticketResultSet() 从对象池中获取AlitripRailTradeCloseticketResultSet
func GetAlitripRailTradeCloseticketResultSet() *AlitripRailTradeCloseticketResultSet {
	return poolAlitripRailTradeCloseticketResultSet.Get().(*AlitripRailTradeCloseticketResultSet)
}

// ReleaseAlitripRailTradeCloseticketResultSet 释放AlitripRailTradeCloseticketResultSet
func ReleaseAlitripRailTradeCloseticketResultSet(v *AlitripRailTradeCloseticketResultSet) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Module = false
	poolAlitripRailTradeCloseticketResultSet.Put(v)
}
