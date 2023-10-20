package alsc

import (
	"sync"
)

// VoucherStatusInfo 结构体
type VoucherStatusInfo struct {
	// 券码
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolVoucherStatusInfo = sync.Pool{
	New: func() any {
		return new(VoucherStatusInfo)
	},
}

// GetVoucherStatusInfo() 从对象池中获取VoucherStatusInfo
func GetVoucherStatusInfo() *VoucherStatusInfo {
	return poolVoucherStatusInfo.Get().(*VoucherStatusInfo)
}

// ReleaseVoucherStatusInfo 释放VoucherStatusInfo
func ReleaseVoucherStatusInfo(v *VoucherStatusInfo) {
	v.VoucherId = ""
	v.Success = false
	poolVoucherStatusInfo.Put(v)
}
