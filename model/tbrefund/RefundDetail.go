package tbrefund

import (
	"sync"
)

// RefundDetail 结构体
type RefundDetail struct {
	// 退款当前可以执行的操作
	AllowedOperations []Operation `json:"allowed_operations,omitempty" xml:"allowed_operations>operation,omitempty"`
	// 退款当前不可以执行的操作
	NotAllowedOperations []Operation `json:"not_allowed_operations,omitempty" xml:"not_allowed_operations>operation,omitempty"`
	// 退款版本号
	RefundVersion int64 `json:"refund_version,omitempty" xml:"refund_version,omitempty"`
}

var poolRefundDetail = sync.Pool{
	New: func() any {
		return new(RefundDetail)
	},
}

// GetRefundDetail() 从对象池中获取RefundDetail
func GetRefundDetail() *RefundDetail {
	return poolRefundDetail.Get().(*RefundDetail)
}

// ReleaseRefundDetail 释放RefundDetail
func ReleaseRefundDetail(v *RefundDetail) {
	v.AllowedOperations = v.AllowedOperations[:0]
	v.NotAllowedOperations = v.NotAllowedOperations[:0]
	v.RefundVersion = 0
	poolRefundDetail.Put(v)
}
