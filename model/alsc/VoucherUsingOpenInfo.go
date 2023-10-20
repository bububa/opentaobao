package alsc

import (
	"sync"
)

// VoucherUsingOpenInfo 结构体
type VoucherUsingOpenInfo struct {
	// 1
	VoucherStatusList []VoucherStatus `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status,omitempty"`
}

var poolVoucherUsingOpenInfo = sync.Pool{
	New: func() any {
		return new(VoucherUsingOpenInfo)
	},
}

// GetVoucherUsingOpenInfo() 从对象池中获取VoucherUsingOpenInfo
func GetVoucherUsingOpenInfo() *VoucherUsingOpenInfo {
	return poolVoucherUsingOpenInfo.Get().(*VoucherUsingOpenInfo)
}

// ReleaseVoucherUsingOpenInfo 释放VoucherUsingOpenInfo
func ReleaseVoucherUsingOpenInfo(v *VoucherUsingOpenInfo) {
	v.VoucherStatusList = v.VoucherStatusList[:0]
	poolVoucherUsingOpenInfo.Put(v)
}
