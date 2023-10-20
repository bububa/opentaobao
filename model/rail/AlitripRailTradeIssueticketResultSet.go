package rail

import (
	"sync"
)

// AlitripRailTradeIssueticketResultSet 结构体
type AlitripRailTradeIssueticketResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功失败
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlitripRailTradeIssueticketResultSet = sync.Pool{
	New: func() any {
		return new(AlitripRailTradeIssueticketResultSet)
	},
}

// GetAlitripRailTradeIssueticketResultSet() 从对象池中获取AlitripRailTradeIssueticketResultSet
func GetAlitripRailTradeIssueticketResultSet() *AlitripRailTradeIssueticketResultSet {
	return poolAlitripRailTradeIssueticketResultSet.Get().(*AlitripRailTradeIssueticketResultSet)
}

// ReleaseAlitripRailTradeIssueticketResultSet 释放AlitripRailTradeIssueticketResultSet
func ReleaseAlitripRailTradeIssueticketResultSet(v *AlitripRailTradeIssueticketResultSet) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Module = false
	poolAlitripRailTradeIssueticketResultSet.Put(v)
}
