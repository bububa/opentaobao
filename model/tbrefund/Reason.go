package tbrefund

import (
	"sync"
)

// Reason 结构体
type Reason struct {
	// 退款原因文案
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 已经影响商品完好
	ReasonTips string `json:"reason_tips,omitempty" xml:"reason_tips,omitempty"`
	// 退款原因ID
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
	v.ReasonTips = ""
	v.ReasonId = 0
	poolReason.Put(v)
}
