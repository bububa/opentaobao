package btrip

import (
	"sync"
)

// MultiRefundPreCalDetail 结构体
type MultiRefundPreCalDetail struct {
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 预计退还金额
	PreRefundMoney int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// 退票手续费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 是否能退票申请
	CanApplyRefund bool `json:"can_apply_refund,omitempty" xml:"can_apply_refund,omitempty"`
}

var poolMultiRefundPreCalDetail = sync.Pool{
	New: func() any {
		return new(MultiRefundPreCalDetail)
	},
}

// GetMultiRefundPreCalDetail() 从对象池中获取MultiRefundPreCalDetail
func GetMultiRefundPreCalDetail() *MultiRefundPreCalDetail {
	return poolMultiRefundPreCalDetail.Get().(*MultiRefundPreCalDetail)
}

// ReleaseMultiRefundPreCalDetail 释放MultiRefundPreCalDetail
func ReleaseMultiRefundPreCalDetail(v *MultiRefundPreCalDetail) {
	v.Name = ""
	v.UserId = ""
	v.PreRefundMoney = 0
	v.RefundFee = 0
	v.CanApplyRefund = false
	poolMultiRefundPreCalDetail.Put(v)
}
