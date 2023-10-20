package alsc

import (
	"sync"
)

// PropertyVerifyOpenInfo 结构体
type PropertyVerifyOpenInfo struct {
	// 券码核销信息
	VoucherStatusList []VoucherStatusInfo `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status_info,omitempty"`
	// 积分核销是否成功
	PointSuccess bool `json:"point_success,omitempty" xml:"point_success,omitempty"`
	// 储值核销是否成功
	ValueSuccess bool `json:"value_success,omitempty" xml:"value_success,omitempty"`
}

var poolPropertyVerifyOpenInfo = sync.Pool{
	New: func() any {
		return new(PropertyVerifyOpenInfo)
	},
}

// GetPropertyVerifyOpenInfo() 从对象池中获取PropertyVerifyOpenInfo
func GetPropertyVerifyOpenInfo() *PropertyVerifyOpenInfo {
	return poolPropertyVerifyOpenInfo.Get().(*PropertyVerifyOpenInfo)
}

// ReleasePropertyVerifyOpenInfo 释放PropertyVerifyOpenInfo
func ReleasePropertyVerifyOpenInfo(v *PropertyVerifyOpenInfo) {
	v.VoucherStatusList = v.VoucherStatusList[:0]
	v.PointSuccess = false
	v.ValueSuccess = false
	poolPropertyVerifyOpenInfo.Put(v)
}
