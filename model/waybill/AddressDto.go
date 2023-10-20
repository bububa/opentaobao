package waybill

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
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
	// 城市，长度小于20
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 详细地址，长度小于256
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区，长度小于20
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 省，长度小于20
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道，长度小于30
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 订购关系id
	WaybillAddressId string `json:"waybill_address_id,omitempty" xml:"waybill_address_id,omitempty"`
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
	v.TownName = ""
	v.AddressDetail = ""
	v.CityName = ""
	v.AreaName = ""
	v.ProvinceName = ""
	v.City = ""
	v.Detail = ""
	v.District = ""
	v.Province = ""
	v.Town = ""
	v.WaybillAddressId = ""
	poolAddressDto.Put(v)
}
