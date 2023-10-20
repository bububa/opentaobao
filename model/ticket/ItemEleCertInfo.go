package ticket

import (
	"sync"
)

// ItemEleCertInfo 结构体
type ItemEleCertInfo struct {
	// 电子凭证 有效期 开始时间
	ExpiryDateStart string `json:"expiry_date_start,omitempty" xml:"expiry_date_start,omitempty"`
	// 电子凭证 有效期 结束时间
	ExpiryDateEnd string `json:"expiry_date_end,omitempty" xml:"expiry_date_end,omitempty"`
	// 门票商品电子凭证信息必填，店铺联系方式
	ShopTel string `json:"shop_tel,omitempty" xml:"shop_tel,omitempty"`
	// 核销服务提供商
	MerchantName string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	// 有效期 过期类型
	ExpiryDateType int64 `json:"expiry_date_type,omitempty" xml:"expiry_date_type,omitempty"`
	// 电子凭证 有效期 天数
	ExpiryDays int64 `json:"expiry_days,omitempty" xml:"expiry_days,omitempty"`
	// 核销门店库id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 售中自动退款比例，0~100
	AutoRefundRate int64 `json:"auto_refund_rate,omitempty" xml:"auto_refund_rate,omitempty"`
	// 过期自动退款比例，0~100
	ExpiredRefundRate int64 `json:"expired_refund_rate,omitempty" xml:"expired_refund_rate,omitempty"`
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
	v.ExpiryDateStart = ""
	v.ExpiryDateEnd = ""
	v.ShopTel = ""
	v.MerchantName = ""
	v.ExpiryDateType = 0
	v.ExpiryDays = 0
	v.PackageId = 0
	v.AutoRefundRate = 0
	v.ExpiredRefundRate = 0
	poolItemEleCertInfo.Put(v)
}
