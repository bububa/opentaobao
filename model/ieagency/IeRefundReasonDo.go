package ieagency

import (
	"sync"
)

// IeRefundReasonDo 结构体
type IeRefundReasonDo struct {
	// 原因描述
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 原因类型
	ReasonType int64 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
}

var poolIeRefundReasonDo = sync.Pool{
	New: func() any {
		return new(IeRefundReasonDo)
	},
}

// GetIeRefundReasonDo() 从对象池中获取IeRefundReasonDo
func GetIeRefundReasonDo() *IeRefundReasonDo {
	return poolIeRefundReasonDo.Get().(*IeRefundReasonDo)
}

// ReleaseIeRefundReasonDo 释放IeRefundReasonDo
func ReleaseIeRefundReasonDo(v *IeRefundReasonDo) {
	v.Reason = ""
	v.ReasonType = 0
	poolIeRefundReasonDo.Put(v)
}
