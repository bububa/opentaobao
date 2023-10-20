package wdk

import (
	"sync"
)

// OfferLicenseInfo 结构体
type OfferLicenseInfo struct {
	// 是否首选证件
	First string `json:"first,omitempty" xml:"first,omitempty"`
	// 证件开始时间
	GmtBegin string `json:"gmt_begin,omitempty" xml:"gmt_begin,omitempty"`
	// 证件结束日期
	GmtValidity string `json:"gmt_validity,omitempty" xml:"gmt_validity,omitempty"`
	// 证件国家
	LicenseCountry string `json:"license_country,omitempty" xml:"license_country,omitempty"`
	// 证件类型
	LicenseType string `json:"license_type,omitempty" xml:"license_type,omitempty"`
	// 证件号码
	LicenseValue string `json:"license_value,omitempty" xml:"license_value,omitempty"`
}

var poolOfferLicenseInfo = sync.Pool{
	New: func() any {
		return new(OfferLicenseInfo)
	},
}

// GetOfferLicenseInfo() 从对象池中获取OfferLicenseInfo
func GetOfferLicenseInfo() *OfferLicenseInfo {
	return poolOfferLicenseInfo.Get().(*OfferLicenseInfo)
}

// ReleaseOfferLicenseInfo 释放OfferLicenseInfo
func ReleaseOfferLicenseInfo(v *OfferLicenseInfo) {
	v.First = ""
	v.GmtBegin = ""
	v.GmtValidity = ""
	v.LicenseCountry = ""
	v.LicenseType = ""
	v.LicenseValue = ""
	poolOfferLicenseInfo.Put(v)
}
