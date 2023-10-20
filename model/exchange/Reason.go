package exchange

import (
	"sync"
)

// Reason 结构体
type Reason struct {
	// 拒绝原因内容
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 拒绝原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}

var poolReason = sync.Pool{
	New: func() any {
		return new(Reason)
	},
}

// GetReason() 从对象池中获取Reason
func GetReason() *Reason {
	return poolReason.Get().(*Reason)
}

// ReleaseReason 释放Reason
func ReleaseReason(v *Reason) {
	v.ReasonText = ""
	v.ReasonId = 0
	poolReason.Put(v)
}
