package alsc

import (
	"sync"
)

// VoucherStatus 结构体
type VoucherStatus struct {
	// 优惠券实例ID/反核销回补ID
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolVoucherStatus = sync.Pool{
	New: func() any {
		return new(VoucherStatus)
	},
}

// GetVoucherStatus() 从对象池中获取VoucherStatus
func GetVoucherStatus() *VoucherStatus {
	return poolVoucherStatus.Get().(*VoucherStatus)
}

// ReleaseVoucherStatus 释放VoucherStatus
func ReleaseVoucherStatus(v *VoucherStatus) {
	v.VoucherId = ""
	v.IsSuccess = false
	poolVoucherStatus.Put(v)
}
