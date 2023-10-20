package aesolution

import (
	"sync"
)

// GlobalAeopTpRefundInfoDto 结构体
type GlobalAeopTpRefundInfoDto struct {
	// refund reason
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// refund status: wait_refund, refund_ok, refund_cancel,  close,  refund_frozen
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// refund type
	RefundType string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// refund time
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// refund coupon amount
	RefundCouponAmt *GlobalMoneyStr `json:"refund_coupon_amt,omitempty" xml:"refund_coupon_amt,omitempty"`
	// refund cash amount
	RefundCashAmt *GlobalMoneyStr `json:"refund_cash_amt,omitempty" xml:"refund_cash_amt,omitempty"`
}

var poolGlobalAeopTpRefundInfoDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpRefundInfoDto)
	},
}

// GetGlobalAeopTpRefundInfoDto() 从对象池中获取GlobalAeopTpRefundInfoDto
func GetGlobalAeopTpRefundInfoDto() *GlobalAeopTpRefundInfoDto {
	return poolGlobalAeopTpRefundInfoDto.Get().(*GlobalAeopTpRefundInfoDto)
}

// ReleaseGlobalAeopTpRefundInfoDto 释放GlobalAeopTpRefundInfoDto
func ReleaseGlobalAeopTpRefundInfoDto(v *GlobalAeopTpRefundInfoDto) {
	v.RefundReason = ""
	v.RefundStatus = ""
	v.RefundType = ""
	v.RefundTime = ""
	v.RefundCouponAmt = nil
	v.RefundCashAmt = nil
	poolGlobalAeopTpRefundInfoDto.Put(v)
}
