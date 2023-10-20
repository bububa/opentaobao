package alsc

import (
	"sync"
)

// ShareMarketingCustomerVoucherInfo 结构体
type ShareMarketingCustomerVoucherInfo struct {
	// 推荐人分享后领取的券（券模板id和对应券实例个数）
	ShareData string `json:"share_data,omitempty" xml:"share_data,omitempty"`
	// 被推荐人领取后，给推荐人的奖励（券模板id和对应券实例个数）
	SharedReceived string `json:"shared_received,omitempty" xml:"shared_received,omitempty"`
}

var poolShareMarketingCustomerVoucherInfo = sync.Pool{
	New: func() any {
		return new(ShareMarketingCustomerVoucherInfo)
	},
}

// GetShareMarketingCustomerVoucherInfo() 从对象池中获取ShareMarketingCustomerVoucherInfo
func GetShareMarketingCustomerVoucherInfo() *ShareMarketingCustomerVoucherInfo {
	return poolShareMarketingCustomerVoucherInfo.Get().(*ShareMarketingCustomerVoucherInfo)
}

// ReleaseShareMarketingCustomerVoucherInfo 释放ShareMarketingCustomerVoucherInfo
func ReleaseShareMarketingCustomerVoucherInfo(v *ShareMarketingCustomerVoucherInfo) {
	v.ShareData = ""
	v.SharedReceived = ""
	poolShareMarketingCustomerVoucherInfo.Put(v)
}
