package alsc

import (
	"sync"
)

// CustomerGetOpenInfo 结构体
type CustomerGetOpenInfo struct {
	// 卡实例信息
	CardOpenInfoList []CardOpenInfo `json:"card_open_info_list,omitempty" xml:"card_open_info_list>card_open_info,omitempty"`
	// 券信息
	VoucherOpenInfoList []VoucherOpenInfo `json:"voucher_open_info_list,omitempty" xml:"voucher_open_info_list>voucher_open_info,omitempty"`
	// 会员信息
	CustomerOpenInfo *CustomerOpenInfo `json:"customer_open_info,omitempty" xml:"customer_open_info,omitempty"`
	// 积分信息
	PointAccountOpenInfo *PointAccountOpenInfo `json:"point_account_open_info,omitempty" xml:"point_account_open_info,omitempty"`
}

var poolCustomerGetOpenInfo = sync.Pool{
	New: func() any {
		return new(CustomerGetOpenInfo)
	},
}

// GetCustomerGetOpenInfo() 从对象池中获取CustomerGetOpenInfo
func GetCustomerGetOpenInfo() *CustomerGetOpenInfo {
	return poolCustomerGetOpenInfo.Get().(*CustomerGetOpenInfo)
}

// ReleaseCustomerGetOpenInfo 释放CustomerGetOpenInfo
func ReleaseCustomerGetOpenInfo(v *CustomerGetOpenInfo) {
	v.CardOpenInfoList = v.CardOpenInfoList[:0]
	v.VoucherOpenInfoList = v.VoucherOpenInfoList[:0]
	v.CustomerOpenInfo = nil
	v.PointAccountOpenInfo = nil
	poolCustomerGetOpenInfo.Put(v)
}
