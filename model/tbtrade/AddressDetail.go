package tbtrade

import (
	"sync"
)

// AddressDetail 结构体
type AddressDetail struct {
	// 详细地址
	DetailedAddress string `json:"detailed_address,omitempty" xml:"detailed_address,omitempty"`
	// 行政区编码
	DivisionCode string `json:"division_code,omitempty" xml:"division_code,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 电话号码
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 国家
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolAddressDetail = sync.Pool{
	New: func() any {
		return new(AddressDetail)
	},
}

// GetAddressDetail() 从对象池中获取AddressDetail
func GetAddressDetail() *AddressDetail {
	return poolAddressDetail.Get().(*AddressDetail)
}

// ReleaseAddressDetail 释放AddressDetail
func ReleaseAddressDetail(v *AddressDetail) {
	v.DetailedAddress = ""
	v.DivisionCode = ""
	v.Area = ""
	v.Telephone = ""
	v.City = ""
	v.Prov = ""
	v.Country = ""
	v.ContactName = ""
	v.BizType = ""
	poolAddressDetail.Put(v)
}
