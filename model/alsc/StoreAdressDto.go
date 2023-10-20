package alsc

import (
	"sync"
)

// StoreAdressDto 结构体
type StoreAdressDto struct {
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 营业面积
	BusinessArea string `json:"business_area,omitempty" xml:"business_area,omitempty"`
	// 维度
	PosY float64 `json:"pos_y,omitempty" xml:"pos_y,omitempty"`
	// 经度
	PosX float64 `json:"pos_x,omitempty" xml:"pos_x,omitempty"`
}

var poolStoreAdressDto = sync.Pool{
	New: func() any {
		return new(StoreAdressDto)
	},
}

// GetStoreAdressDto() 从对象池中获取StoreAdressDto
func GetStoreAdressDto() *StoreAdressDto {
	return poolStoreAdressDto.Get().(*StoreAdressDto)
}

// ReleaseStoreAdressDto 释放StoreAdressDto
func ReleaseStoreAdressDto(v *StoreAdressDto) {
	v.DetailAddress = ""
	v.Town = ""
	v.Area = ""
	v.City = ""
	v.Province = ""
	v.BusinessArea = ""
	v.PosY = 0
	v.PosX = 0
	poolStoreAdressDto.Put(v)
}
