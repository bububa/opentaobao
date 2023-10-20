package logistic

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
	// first name of receiver
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name of receiver
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// receiver&#39;s city
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// receiver&#39;s federal_tax_id
	FederalTaxId string `json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// receiver&#39;s country
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// zip code of ship to place
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// receiver&#39;s State
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// receiver&#39;s district and street
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// receiver&#39;s street number
	AddressNumber string `json:"address_number,omitempty" xml:"address_number,omitempty"`
	// email of receiver
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// cell phone of receiver
	CellPhone string `json:"cell_phone,omitempty" xml:"cell_phone,omitempty"`
	// shipping additional
	Additional string `json:"additional,omitempty" xml:"additional,omitempty"`
	// 镇/街道
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 区级地址
	County string `json:"county,omitempty" xml:"county,omitempty"`
	// 省级地址
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道地址
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

var poolAddressDto = sync.Pool{
	New: func() any {
		return new(AddressDto)
	},
}

// GetAddressDto() 从对象池中获取AddressDto
func GetAddressDto() *AddressDto {
	return poolAddressDto.Get().(*AddressDto)
}

// ReleaseAddressDto 释放AddressDto
func ReleaseAddressDto(v *AddressDto) {
	v.FirstName = ""
	v.LastName = ""
	v.City = ""
	v.FederalTaxId = ""
	v.Country = ""
	v.ZipCode = ""
	v.State = ""
	v.Address = ""
	v.AddressNumber = ""
	v.Email = ""
	v.CellPhone = ""
	v.Additional = ""
	v.TownName = ""
	v.AddressDetail = ""
	v.CityName = ""
	v.AreaName = ""
	v.ProvinceName = ""
	v.County = ""
	v.Province = ""
	v.Town = ""
	v.DetailAddress = ""
	poolAddressDto.Put(v)
}
