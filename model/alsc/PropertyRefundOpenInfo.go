package alsc

import (
	"sync"
)

// PropertyRefundOpenInfo 结构体
type PropertyRefundOpenInfo struct {
	// 券实例状态
	VoucherStatusList []VoucherStatusInfo `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status_info,omitempty"`
	// 回退积分是否成功
	PointSuccess bool `json:"point_success,omitempty" xml:"point_success,omitempty"`
	// 回退储值是否成功
	ValueSuccess bool `json:"value_success,omitempty" xml:"value_success,omitempty"`
}

var poolPropertyRefundOpenInfo = sync.Pool{
	New: func() any {
		return new(PropertyRefundOpenInfo)
	},
}

// GetPropertyRefundOpenInfo() 从对象池中获取PropertyRefundOpenInfo
func GetPropertyRefundOpenInfo() *PropertyRefundOpenInfo {
	return poolPropertyRefundOpenInfo.Get().(*PropertyRefundOpenInfo)
}

// ReleasePropertyRefundOpenInfo 释放PropertyRefundOpenInfo
func ReleasePropertyRefundOpenInfo(v *PropertyRefundOpenInfo) {
	v.VoucherStatusList = v.VoucherStatusList[:0]
	v.PointSuccess = false
	v.ValueSuccess = false
	poolPropertyRefundOpenInfo.Put(v)
}
