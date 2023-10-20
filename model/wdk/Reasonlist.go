package wdk

import (
	"sync"
)

// Reasonlist 结构体
type Reasonlist struct {
	// 退款原因说明
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 退款原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}

var poolReasonlist = sync.Pool{
	New: func() any {
		return new(Reasonlist)
	},
}

// GetReasonlist() 从对象池中获取Reasonlist
func GetReasonlist() *Reasonlist {
	return poolReasonlist.Get().(*Reasonlist)
}

// ReleaseReasonlist 释放Reasonlist
func ReleaseReasonlist(v *Reasonlist) {
	v.ReasonText = ""
	v.ReasonId = 0
	poolReasonlist.Put(v)
}
