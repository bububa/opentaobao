package travel

import (
	"sync"
)

// PontusTravelItemEleCertInfo 结构体
type PontusTravelItemEleCertInfo struct {
	// 电子凭证 有效期 开始时间
	ExpiryDateStart string `json:"expiry_date_start,omitempty" xml:"expiry_date_start,omitempty"`
	// 电子凭证 有效期 结束时间
	ExpiryDateEnd string `json:"expiry_date_end,omitempty" xml:"expiry_date_end,omitempty"`
	// 电子凭证有效期 过期类型。1：xxxx-xx-xx 到 xxxx-xx-xx； 2：购买成功日 至 xxxx-xx-xx； 3：购买成功 xx 天内有效
	ExpiryDateType int64 `json:"expiry_date_type,omitempty" xml:"expiry_date_type,omitempty"`
	// 电子凭证 有效期 天数
	ExpiryDays int64 `json:"expiry_days,omitempty" xml:"expiry_days,omitempty"`
	// 核销门店库id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 售中自动退款比例
	AutoRefundRate int64 `json:"auto_refund_rate,omitempty" xml:"auto_refund_rate,omitempty"`
	// 过期自动退款比例
	ExpiredRefundRate int64 `json:"expired_refund_rate,omitempty" xml:"expired_refund_rate,omitempty"`
}

var poolPontusTravelItemEleCertInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemEleCertInfo)
	},
}

// GetPontusTravelItemEleCertInfo() 从对象池中获取PontusTravelItemEleCertInfo
func GetPontusTravelItemEleCertInfo() *PontusTravelItemEleCertInfo {
	return poolPontusTravelItemEleCertInfo.Get().(*PontusTravelItemEleCertInfo)
}

// ReleasePontusTravelItemEleCertInfo 释放PontusTravelItemEleCertInfo
func ReleasePontusTravelItemEleCertInfo(v *PontusTravelItemEleCertInfo) {
	v.ExpiryDateStart = ""
	v.ExpiryDateEnd = ""
	v.ExpiryDateType = 0
	v.ExpiryDays = 0
	v.PackageId = 0
	v.AutoRefundRate = 0
	v.ExpiredRefundRate = 0
	poolPontusTravelItemEleCertInfo.Put(v)
}
