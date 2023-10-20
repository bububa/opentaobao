package tmallservice

import (
	"sync"
)

// RefundWorkCardTo 结构体
type RefundWorkCardTo struct {
	// failureList
	FailureList []int64 `json:"failure_list,omitempty" xml:"failure_list>int64,omitempty"`
	// notExistingList
	NotExistingList []int64 `json:"not_existing_list,omitempty" xml:"not_existing_list>int64,omitempty"`
	// refundList
	RefundList []int64 `json:"refund_list,omitempty" xml:"refund_list>int64,omitempty"`
	// existingFailure
	ExistingFailure bool `json:"existing_failure,omitempty" xml:"existing_failure,omitempty"`
}

var poolRefundWorkCardTo = sync.Pool{
	New: func() any {
		return new(RefundWorkCardTo)
	},
}

// GetRefundWorkCardTo() 从对象池中获取RefundWorkCardTo
func GetRefundWorkCardTo() *RefundWorkCardTo {
	return poolRefundWorkCardTo.Get().(*RefundWorkCardTo)
}

// ReleaseRefundWorkCardTo 释放RefundWorkCardTo
func ReleaseRefundWorkCardTo(v *RefundWorkCardTo) {
	v.FailureList = v.FailureList[:0]
	v.NotExistingList = v.NotExistingList[:0]
	v.RefundList = v.RefundList[:0]
	v.ExistingFailure = false
	poolRefundWorkCardTo.Put(v)
}
