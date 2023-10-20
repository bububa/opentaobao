package wlbimports

import (
	"sync"
)

// PickupInfo 结构体
type PickupInfo struct {
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 公司（非自寄模式必填）
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 资源名称
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 资源code（非自寄模式必填）
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 电话（非自寄模式必填）
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 名称（非自寄模式必填）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 详细地址（最大135个字符）（非自寄模式必填）
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 邮箱（非自寄模式必填）
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国家id（非自寄模式必填）
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 菜鸟地址id（非自寄模式必填）
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolPickupInfo = sync.Pool{
	New: func() any {
		return new(PickupInfo)
	},
}

// GetPickupInfo() 从对象池中获取PickupInfo
func GetPickupInfo() *PickupInfo {
	return poolPickupInfo.Get().(*PickupInfo)
}

// ReleasePickupInfo 释放PickupInfo
func ReleasePickupInfo(v *PickupInfo) {
	v.Country = ""
	v.ZipCode = ""
	v.City = ""
	v.CompanyName = ""
	v.ResourceName = ""
	v.ResourceCode = ""
	v.Province = ""
	v.Phone = ""
	v.Name = ""
	v.DetailAddress = ""
	v.Email = ""
	v.CountryId = 0
	v.AreaId = 0
	poolPickupInfo.Put(v)
}
