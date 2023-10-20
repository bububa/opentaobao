package ascpchannel

import (
	"sync"
)

// Closerefundorderrequest 结构体
type Closerefundorderrequest struct {
	// 退款单号
	RefundNo string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
	// 外部退款单号
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
}

var poolCloserefundorderrequest = sync.Pool{
	New: func() any {
		return new(Closerefundorderrequest)
	},
}

// GetCloserefundorderrequest() 从对象池中获取Closerefundorderrequest
func GetCloserefundorderrequest() *Closerefundorderrequest {
	return poolCloserefundorderrequest.Get().(*Closerefundorderrequest)
}

// ReleaseCloserefundorderrequest 释放Closerefundorderrequest
func ReleaseCloserefundorderrequest(v *Closerefundorderrequest) {
	v.RefundNo = ""
	v.OutRefundNo = ""
	poolCloserefundorderrequest.Put(v)
}
