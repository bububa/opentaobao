package travel

import (
	"sync"
)

// ItemEleCertInfo 结构体
type ItemEleCertInfo struct {
	// 殊必填（expiryDateType为1或2时必填），电子凭证 有效期 结束时间
	ExpiryDateEnd string `json:"expiry_date_end,omitempty" xml:"expiry_date_end,omitempty"`
	// 特殊必填（expiryDateType为1时必填），电子凭证 有效期 开始时间
	ExpiryDateStart string `json:"expiry_date_start,omitempty" xml:"expiry_date_start,omitempty"`
	// 可选，售中自动退款比例，0~100。默认为0，即不支持售中自动退款；当值为1~100时表示售中自动退款的比例
	AutoRefundRate int64 `json:"auto_refund_rate,omitempty" xml:"auto_refund_rate,omitempty"`
	// 可选，过期自动退款比例，0~100。默认为0，即不支持过期自动退款比例；当值为1~100时表示过期自动退款的比例
	ExpiredRefundRate int64 `json:"expired_refund_rate,omitempty" xml:"expired_refund_rate,omitempty"`
	// 必填，电子凭证有效期 过期类型。1：xxxx-xx-xx 到 xxxx-xx-xx； 2：购买成功日 至 xxxx-xx-xx； 3：购买成功 xx 天内有效
	ExpiryDateType int64 `json:"expiry_date_type,omitempty" xml:"expiry_date_type,omitempty"`
	// 特殊必填（expiryDateType为3时必填），电子凭证 有效期 天数
	ExpiryDays int64 `json:"expiry_days,omitempty" xml:"expiry_days,omitempty"`
	// 必填，核销门店库id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
}

var poolItemEleCertInfo = sync.Pool{
	New: func() any {
		return new(ItemEleCertInfo)
	},
}

// GetItemEleCertInfo() 从对象池中获取ItemEleCertInfo
func GetItemEleCertInfo() *ItemEleCertInfo {
	return poolItemEleCertInfo.Get().(*ItemEleCertInfo)
}

// ReleaseItemEleCertInfo 释放ItemEleCertInfo
func ReleaseItemEleCertInfo(v *ItemEleCertInfo) {
	v.ExpiryDateEnd = ""
	v.ExpiryDateStart = ""
	v.AutoRefundRate = 0
	v.ExpiredRefundRate = 0
	v.ExpiryDateType = 0
	v.ExpiryDays = 0
	v.PackageId = 0
	poolItemEleCertInfo.Put(v)
}
