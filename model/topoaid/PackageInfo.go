package topoaid

import (
	"sync"
)

// PackageInfo 结构体
type PackageInfo struct {
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 订单隐私号
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 隐私号过期时间
	SecretNoExpireTime string `json:"secret_no_expire_time,omitempty" xml:"secret_no_expire_time,omitempty"`
	// 收件人手机号,如果用户未授权手机号则不返回
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 快递公司代码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 代收点类型
	StationType string `json:"station_type,omitempty" xml:"station_type,omitempty"`
	// 是否为隐私订单
	PrivacyOrder bool `json:"privacy_order,omitempty" xml:"privacy_order,omitempty"`
}

var poolPackageInfo = sync.Pool{
	New: func() any {
		return new(PackageInfo)
	},
}

// GetPackageInfo() 从对象池中获取PackageInfo
func GetPackageInfo() *PackageInfo {
	return poolPackageInfo.Get().(*PackageInfo)
}

// ReleasePackageInfo 释放PackageInfo
func ReleasePackageInfo(v *PackageInfo) {
	v.MailNo = ""
	v.SecretNo = ""
	v.SecretNoExpireTime = ""
	v.ReceiverMobile = ""
	v.CpCode = ""
	v.StationType = ""
	v.PrivacyOrder = false
	poolPackageInfo.Put(v)
}
