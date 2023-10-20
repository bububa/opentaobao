package ticket

import (
	"sync"
)

// AlitripTicketRuleUploadResultSet 结构体
type AlitripTicketRuleUploadResultSet struct {
	// 规则维护结果
	FirstResult *TopTicketRuleResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

var poolAlitripTicketRuleUploadResultSet = sync.Pool{
	New: func() any {
		return new(AlitripTicketRuleUploadResultSet)
	},
}

// GetAlitripTicketRuleUploadResultSet() 从对象池中获取AlitripTicketRuleUploadResultSet
func GetAlitripTicketRuleUploadResultSet() *AlitripTicketRuleUploadResultSet {
	return poolAlitripTicketRuleUploadResultSet.Get().(*AlitripTicketRuleUploadResultSet)
}

// ReleaseAlitripTicketRuleUploadResultSet 释放AlitripTicketRuleUploadResultSet
func ReleaseAlitripTicketRuleUploadResultSet(v *AlitripTicketRuleUploadResultSet) {
	v.FirstResult = nil
	poolAlitripTicketRuleUploadResultSet.Put(v)
}
