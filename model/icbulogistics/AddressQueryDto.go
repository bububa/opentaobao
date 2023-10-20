package icbulogistics

import (
	"sync"
)

// AddressQueryDto 结构体
type AddressQueryDto struct {
	// 国家code
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 查询关键词
	SearchText string `json:"search_text,omitempty" xml:"search_text,omitempty"`
	// 省ID
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 是否包含子节点
	WithChildren bool `json:"with_children,omitempty" xml:"with_children,omitempty"`
}

var poolAddressQueryDto = sync.Pool{
	New: func() any {
		return new(AddressQueryDto)
	},
}

// GetAddressQueryDto() 从对象池中获取AddressQueryDto
func GetAddressQueryDto() *AddressQueryDto {
	return poolAddressQueryDto.Get().(*AddressQueryDto)
}

// ReleaseAddressQueryDto 释放AddressQueryDto
func ReleaseAddressQueryDto(v *AddressQueryDto) {
	v.CountryCode = ""
	v.SearchText = ""
	v.ProvinceId = 0
	v.CityId = 0
	v.WithChildren = false
	poolAddressQueryDto.Put(v)
}
