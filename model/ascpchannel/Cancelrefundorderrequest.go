package ascpchannel

import (
	"sync"
)

// Cancelrefundorderrequest 结构体
type Cancelrefundorderrequest struct {
	// 退款单号
	RefundNo string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
	// 外部退款单号
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
}

var poolCancelrefundorderrequest = sync.Pool{
	New: func() any {
		return new(Cancelrefundorderrequest)
	},
}

// GetCancelrefundorderrequest() 从对象池中获取Cancelrefundorderrequest
func GetCancelrefundorderrequest() *Cancelrefundorderrequest {
	return poolCancelrefundorderrequest.Get().(*Cancelrefundorderrequest)
}

// ReleaseCancelrefundorderrequest 释放Cancelrefundorderrequest
func ReleaseCancelrefundorderrequest(v *Cancelrefundorderrequest) {
	v.RefundNo = ""
	v.OutRefundNo = ""
	poolCancelrefundorderrequest.Put(v)
}
