package topoaid

import (
	"sync"
)

// IsPrivacyPackageRequest 结构体
type IsPrivacyPackageRequest struct {
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 柜机黑名单手机号前7位，如有多个用英文逗号,分隔，每个元素仅前七位有效
	BlackMobiles string `json:"black_mobiles,omitempty" xml:"black_mobiles,omitempty"`
}

var poolIsPrivacyPackageRequest = sync.Pool{
	New: func() any {
		return new(IsPrivacyPackageRequest)
	},
}

// GetIsPrivacyPackageRequest() 从对象池中获取IsPrivacyPackageRequest
func GetIsPrivacyPackageRequest() *IsPrivacyPackageRequest {
	return poolIsPrivacyPackageRequest.Get().(*IsPrivacyPackageRequest)
}

// ReleaseIsPrivacyPackageRequest 释放IsPrivacyPackageRequest
func ReleaseIsPrivacyPackageRequest(v *IsPrivacyPackageRequest) {
	v.MailNo = ""
	v.CpCode = ""
	v.BlackMobiles = ""
	poolIsPrivacyPackageRequest.Put(v)
}
