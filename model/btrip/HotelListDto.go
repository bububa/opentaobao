package btrip

import (
	"sync"
)

// HotelListDto 结构体
type HotelListDto struct {
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 供应商code
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 最低价，单位分
	LowPrice int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
	// 酒店标准ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 协议价标识
	IsProtocol bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
}

var poolHotelListDto = sync.Pool{
	New: func() any {
		return new(HotelListDto)
	},
}

// GetHotelListDto() 从对象池中获取HotelListDto
func GetHotelListDto() *HotelListDto {
	return poolHotelListDto.Get().(*HotelListDto)
}

// ReleaseHotelListDto 释放HotelListDto
func ReleaseHotelListDto(v *HotelListDto) {
	v.Name = ""
	v.SupplierCode = ""
	v.SupplierName = ""
	v.LowPrice = 0
	v.Shid = 0
	v.IsProtocol = false
	poolHotelListDto.Put(v)
}
