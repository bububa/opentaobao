package ascp

import (
	"sync"
)

// OrderReceiverPrivacyData 结构体
type OrderReceiverPrivacyData struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 移动电话隐私
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 收件人所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收件人所在区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 收件人所在街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}

var poolOrderReceiverPrivacyData = sync.Pool{
	New: func() any {
		return new(OrderReceiverPrivacyData)
	},
}

// GetOrderReceiverPrivacyData() 从对象池中获取OrderReceiverPrivacyData
func GetOrderReceiverPrivacyData() *OrderReceiverPrivacyData {
	return poolOrderReceiverPrivacyData.Get().(*OrderReceiverPrivacyData)
}

// ReleaseOrderReceiverPrivacyData 释放OrderReceiverPrivacyData
func ReleaseOrderReceiverPrivacyData(v *OrderReceiverPrivacyData) {
	v.Name = ""
	v.Phone = ""
	v.CountryCode = ""
	v.Province = ""
	v.DetailAddress = ""
	v.City = ""
	v.District = ""
	v.Town = ""
	poolOrderReceiverPrivacyData.Put(v)
}
